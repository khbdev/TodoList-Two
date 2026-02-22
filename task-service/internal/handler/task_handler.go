package handler

import (
	"context"
	"strconv"
	"time"

	"task-service/internal/domain"
	"task-service/pkg"

	taskpb "github.com/khbdev/todolist-proto/proto/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type TaskHandler struct {
	taskpb.UnimplementedTaskServiceServer
	uc domain.ReminderUsecase
}

func NewTaskHandler(uc domain.ReminderUsecase) *TaskHandler {
	return &TaskHandler{uc: uc}
}



func userIDFromCtx(ctx context.Context) (uint, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, status.Error(codes.Unauthenticated, "missing metadata")
	}

	vals := md.Get("user_id")
	if len(vals) == 0 {
		return 0, status.Error(codes.Unauthenticated, "missing user_id")
	}

	u64, err := strconv.ParseUint(vals[0], 10, 64)
	if err != nil || u64 == 0 {
		return 0, status.Error(codes.Unauthenticated, "invalid user_id")
	}

	return uint(u64), nil
}

func parseRemindAt(s string) (time.Time, error) {

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, status.Error(codes.InvalidArgument, "invalid remind_at (RFC3339 required)")
	}
	return t, nil
}



func (h *TaskHandler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	remindAt, err := parseRemindAt(req.GetRemindAt())
	if err != nil {
		return nil, err
	}

	t, err := h.uc.Create(ctx, uint(req.GetUserId()), req.GetTitle(), remindAt)
	if err != nil {
		return nil, pkg.MapError(err)
	}

	return &taskpb.CreateTaskResponse{
		Task: pkg.ToProtoTask(t),
	}, nil
}

func (h *TaskHandler) GetTasksByUser(ctx context.Context, req *taskpb.GetTasksByUserRequest) (*taskpb.GetTasksByUserResponse, error) {
	list, err := h.uc.GetByUser(ctx, uint(req.GetUserId()))
	if err != nil {
		return nil, pkg.MapError(err)
	}

	return &taskpb.GetTasksByUserResponse{
		Tasks: pkg.ToProtoTasks(list),
	}, nil
}

func (h *TaskHandler) GetTaskById(ctx context.Context, req *taskpb.GetTaskByIdRequest) (*taskpb.GetTaskByIdResponse, error) {

	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	t, err := h.uc.GetByID(ctx, userID, uint(req.GetTaskId()))
	if err != nil {
		return nil, pkg.MapError(err)
	}

	return &taskpb.GetTaskByIdResponse{
		Task: pkg.ToProtoTask(t),
	}, nil
}

func (h *TaskHandler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	remindAt, err := parseRemindAt(req.GetRemindAt())
	if err != nil {
		return nil, err
	}

	
	updated, err := h.uc.Update(ctx, userID, uint(req.GetTaskId()), req.GetTitle(), remindAt, false)
	if err != nil {
		return nil, pkg.MapError(err)
	}

	return &taskpb.UpdateTaskResponse{
		Task:    pkg.ToProtoTask(updated),
		Success: true,
	}, nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if err := h.uc.Delete(ctx, userID, uint(req.GetTaskId())); err != nil {
		return nil, pkg.MapError(err)
	}

	return &taskpb.DeleteTaskResponse{Success: true}, nil
}