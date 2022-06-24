import { useQuery } from "react-query";
import { getPosts } from "../services/posts.service";

export const usePostsQuery = () => {
  return useQuery("posts", getPosts);
};
