package service

import (
	"beer-shop/app/catalog/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "beer-shop/api/catalog/service/v1"
)

type CatalogService struct {
	pb.UnimplementedCatalogServer

	bc  *biz.BeerUseCase
	log *log.Helper
}

func NewCatalogService(bc *biz.BeerUseCase, logger log.Logger) *CatalogService {
	return &CatalogService{

		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}

func (s *CatalogService) CreateBeer(ctx context.Context, req *pb.CreateBeerReq) (*pb.CreateBeerReply, error) {
	b := &biz.Beer{
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Create(ctx, b)
	img := make([]*pb.CreateBeerReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &pb.CreateBeerReply_Image{Url: i.URL})
	}
	return &pb.CreateBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) GetBeer(ctx context.Context, req *pb.GetBeerReq) (*pb.GetBeerReply, error) {
	x, err := s.bc.Get(ctx, req.Id)
	img := make([]*pb.GetBeerReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &pb.GetBeerReply_Image{Url: i.URL})
	}
	return &pb.GetBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) UpdateBeer(ctx context.Context, req *pb.UpdateBeerReq) (*pb.UpdateBeerReply, error) {
	b := &biz.Beer{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Update(ctx, b)
	img := make([]*pb.UpdateBeerReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &pb.UpdateBeerReply_Image{Url: i.URL})
	}
	return &pb.UpdateBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) ListBeer(ctx context.Context, req *pb.ListBeerReq) (*pb.ListBeerReply, error) {
	rv, err := s.bc.List(ctx, req.PageNum, req.PageSize)
	rs := make([]*pb.ListBeerReply_Beer, 0)
	for _, x := range rv {
		img := make([]*pb.ListBeerReply_Beer_Image, 0)
		for _, i := range x.Images {
			img = append(img, &pb.ListBeerReply_Beer_Image{Url: i.URL})
		}
		rs = append(rs, &pb.ListBeerReply_Beer{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Image:       img,
		})
	}
	return &pb.ListBeerReply{
		Results: rs,
	}, err
}

func (s *CatalogService) ListBeerNextToken(ctx context.Context, req *pb.ListBeerNextTokenReq) (*pb.ListBeerReplyNextToken, error) {
	rv, token, err := s.bc.ListNext(ctx, req.PageSize, req.PageToken)
	rs := make([]*pb.ListBeerReplyNextToken_Beer, 0)
	for _, x := range rv {
		img := make([]*pb.ListBeerReplyNextToken_Beer_Image, 0)
		for _, i := range x.Images {
			img = append(img, &pb.ListBeerReplyNextToken_Beer_Image{Url: i.URL})
		}
		rs = append(rs, &pb.ListBeerReplyNextToken_Beer{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Image:       img,
		})
	}
	return &pb.ListBeerReplyNextToken{
		Results:       rs,
		NextPageToken: token,
	}, err
}
