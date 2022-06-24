import { useQuery } from "react-query";
import { getReplies } from "../services/replies.service";

export const useRepliesQuery = (postId: string) => {
  return useQuery(`Replies-${postId}`, () => getReplies({ post_id: postId }));
};
