import { useMutation } from "react-query";
import { queryClient } from "../App";
import { deleteReply } from "../services/replies.service";

export const useDeleteReplyMutation = (postId: string) => {
  return useMutation(deleteReply, {
    onSuccess: () => {
      queryClient.refetchQueries(`Replies-${postId}`);
    },
  });
};
