import { useMutation } from "react-query";
import { queryClient } from "../App";
import { updateReply } from "../services/replies.service";

export const useEditReplyMutation = (postId: string) => {
  return useMutation(updateReply, {
    onSuccess: () => {
      queryClient.refetchQueries(`Replies-${postId}`);
    },
  });
};
