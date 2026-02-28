package handler

import (
	"context"

	"notes-service/internal/domain"
	"notes-service/pkg"

	notespb "github.com/khbdev/todolist-proto/proto/notes"
)

type NotesHandler struct {
	notespb.UnimplementedNotesServiceServer
	uc domain.TodoUsecase
}

func NewNotesHandler(uc domain.TodoUsecase) *NotesHandler {
	return &NotesHandler{uc: uc}
}

func (h *NotesHandler) CreateNote(ctx context.Context, req *notespb.CreateNoteRequest) (*notespb.NoteResponse, error) {
	t, err := h.uc.Create(ctx, uint(req.GetUserId()), req.GetTitle(), req.GetText())
	if err != nil {
		return nil, pkg.MapError(err)
	}
	return &notespb.NoteResponse{Note: pkg.ToProtoNote(t)}, nil
}

func (h *NotesHandler) GetNotesByUser(ctx context.Context, req *notespb.GetNotesByUserRequest) (*notespb.NotesListResponse, error) {
	list, err := h.uc.GetAll(ctx, uint(req.GetUserId()))
	if err != nil {
		return nil, pkg.MapError(err)
	}
	return &notespb.NotesListResponse{Notes: pkg.ToProtoNotes(list)}, nil
}

func (h *NotesHandler) GetNoteByID(ctx context.Context, req *notespb.GetNoteByIDRequest) (*notespb.NoteResponse, error) {
	t, err := h.uc.GetByID(ctx, uint(req.GetUserId()), uint(req.GetNoteId()))
	if err != nil {
		return nil, pkg.MapError(err)
	}
	return &notespb.NoteResponse{Note: pkg.ToProtoNote(t)}, nil
}

func (h *NotesHandler) UpdateNote(ctx context.Context, req *notespb.UpdateNoteRequest) (*notespb.NoteResponse, error) {
	t, err := h.uc.Update(ctx, uint(req.GetUserId()), uint(req.GetNoteId()), req.GetTitle(), req.GetText())
	if err != nil {
		return nil, pkg.MapError(err)
	}
	return &notespb.NoteResponse{Note: pkg.ToProtoNote(t)}, nil
}

func (h *NotesHandler) DeleteNote(ctx context.Context, req *notespb.DeleteNoteRequest) (*notespb.DeleteNoteResponse, error) {
	if err := h.uc.Delete(ctx, uint(req.GetUserId()), uint(req.GetNoteId())); err != nil {
		return nil, pkg.MapError(err)
	}
	return &notespb.DeleteNoteResponse{Success: true}, nil
}