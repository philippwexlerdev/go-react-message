import { useMutation } from "react-query";
import { queryClient } from "../App";
import { editPost } from "../services/posts.service";

export const useEditPostMutation = () => {
  return useMutation(editPost, {
    onSuccess: () => {
      queryClient.refetchQueries("posts");
    },
  });
};
