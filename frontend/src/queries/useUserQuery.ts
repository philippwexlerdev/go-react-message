import { useQuery } from "react-query";
import { getUser } from "../services/user.service";

export const useUserQuery = (userId: string) => {
  return useQuery(`User-${userId}`, () => getUser({ user_id: userId }));
};
