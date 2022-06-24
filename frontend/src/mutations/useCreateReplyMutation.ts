import { useMutation } from "react-query";
import { queryClient } from "../App";
import { createReply } from "../services/replies.service";

export const useCreatReplyMutation = (postId: string) => {
  return useMutation(createReply, {
    onSuccess: () => {
      queryClient.refetchQueries(`Replies-${postId}`);
    },
  });
};
