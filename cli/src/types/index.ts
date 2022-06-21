export interface ITodoItem {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string;
  Title: string;
  Description: string;
  CompletedAt?: string;
  UserID: number;
}
