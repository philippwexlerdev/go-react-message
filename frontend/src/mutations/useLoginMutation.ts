import { useMutation } from "react-query";
import { loginUser } from "../services/user.service";

export const useLoginMutation = () => {
  return useMutation(loginUser);
};
