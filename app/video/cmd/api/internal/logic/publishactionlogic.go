package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stringx"
	"go_code/Doul/app/video/cmd/rpc/pb"
	"go_code/Doul/common/tool"
	"go_code/Doul/common/xerr"
	"mime/multipart"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"go_code/Doul/app/video/cmd/api/internal/svc"
	"go_code/Doul/app/video/cmd/api/internal/types"
)

const (
	defaultMultipartMemory = 32 << 10 // 32 MB
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(r *http.Request) (resp *types.PublishActionResp, err error) {
	// step1 : parse the multipart form and check the video format
	if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
		return nil, err
	}
	title := r.MultipartForm.Value["title"][0]
	_, handler, err := r.FormFile("data")

	if isValid, err := l.CheckValid(handler); !isValid || err != nil {
		return nil, err
	}

	if err != nil || stringx.HasEmpty(title) {
		if stringx.HasEmpty(title) {
			return nil, errors.New("empty title")
		}
		return nil, err
	}

	fd, err := handler.Open()
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	// step2 : save to the oss
	videoUrl, coverUrl, err := l.SaveVideoToOSS(handler.Filename, fd)
	if err != nil {
		return nil, err
	}

	// step3 : insert the record
	if _, err = l.svcCtx.VideoRpc.PublishAction(l.ctx, &pb.PublishReq{
		Title:    title,
		CoverUrl: coverUrl,
		VideoUrl: videoUrl,
		UserId:   tool.GetUidFromCtx(l.ctx),
	}); err != nil {
		return nil, err
	}
	logx.Infof("user %d upload the new video %s", tool.GetUidFromCtx(l.ctx), title)
	return &types.PublishActionResp{}, nil
}

// CheckValid 检查视频格式是否正确
func (l *PublishActionLogic) CheckValid(handler *multipart.FileHeader) (isValid bool, err error) {
	//最多16M视频
	if handler.Size > 16*1024*1024 {
		return false, xerr.NewErrMsg("视频文件大小超出限制！")
	}
	var fileName = handler.Filename
	if fileName[len(fileName)-4:] != ".mp4" {
		return false, xerr.NewErrMsg("上传文件格式不正确！")
	}
	return true, nil
}

// SaveVideoToOSS  save the video to the oss
func (l *PublishActionLogic) SaveVideoToOSS(videoName string, fd multipart.File) (videoUrl string, coverUrl string, err error) {
	//上传 参数1为上传地址 参数2为文件句柄
	// doc : https://www.alibabacloud.com/help/zh/object-storage-service/latest/simple-upload-5
	err = l.svcCtx.OSSEngine.PutObject(l.svcCtx.Config.OSSConf.TargetPath+videoName, fd)
	if err != nil {
		return "", "", err
	}
	//回传视频和封面地址
	videoUrl = fmt.Sprintf("%s/%s%s", l.svcCtx.Config.OSSConf.TargetURL, l.svcCtx.Config.OSSConf.TargetPath, videoName)
	coverUrl = fmt.Sprintf("https://kauizhaotan.oss-cn-shanghai.aliyuncs.com/%s%s"+
		"?x-oss-process=video/snapshot,t_500,m_fast", l.svcCtx.Config.OSSConf.TargetPath, videoName)
	return videoUrl, coverUrl, nil
}
