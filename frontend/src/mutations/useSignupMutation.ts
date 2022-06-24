import { useMutation } from "react-query";
import { signupUser } from "../services/user.service";

export const useSignupMutation = () => {
  return useMutation(signupUser);
};
