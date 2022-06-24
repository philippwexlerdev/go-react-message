// import React, { createContext, useState } from "react";

// interface AuthContextInterface {
//   token?: string;
//   setToken?: (token: string) => void;
// }
// interface AuthContextProviderProps {
//   children: React.ReactNode[] | React.ReactNode;
// }

// export const AuthContext = createContext<AuthContextInterface | null>(null);

// export const AuthContextProvider: React.FC<React.ReactNode> = ({
//   children,
// }: AuthContextProviderProps) => {
//   const [token, setToken] = useState("");
//   return (
//     <AuthContext.Provider value={{ token, setToken }}>
//       {children}
//     </AuthContext.Provider>
//   );
// };

export {};
