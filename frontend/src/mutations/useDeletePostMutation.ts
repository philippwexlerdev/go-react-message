import { useMutation } from "react-query";
import { queryClient } from "../App";
import { deletePost } from "../services/posts.service";

export const useDeletePostMutation = () => {
  return useMutation(deletePost, {
    onSuccess: () => {
      queryClient.refetchQueries("posts");
    },
  });
};
