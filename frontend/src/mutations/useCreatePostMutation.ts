import { useMutation } from "react-query";
import { queryClient } from "../App";
import { createPost } from "../services/posts.service";

export const useCreatePostMutation = () => {
  return useMutation(createPost, {
    onSuccess: () => {
      queryClient.refetchQueries("posts");
    },
  });
};
