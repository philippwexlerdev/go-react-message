import React from "react";
import { Routes, Route } from "react-router-dom";
import AuthLayout from "../layout/AuthLayout";

import NoAuthLayout from "../layout/NoAuthLayout";
import Home from "../pages/Home/Home";
import Login from "../pages/Login/Login";
import Signup from "../pages/Signup/Signup";

const allRoutes = [
  {
    path: "/",
    Component: Home,
    Layout: AuthLayout,
  },
  {
    path: "/login",
    Component: Login,
    Layout: NoAuthLayout,
  },
  {
    path: "/signup",
    Component: Signup,
    Layout: NoAuthLayout,
  },
];

export default function AllRoutes() {
  return (
    <Routes>
      {allRoutes.map(({ path, Component, Layout }) => (
        <Route
          key={path}
          path={path}
          element={
            <Layout>
              <Component />
            </Layout>
          }
        />
      ))}
    </Routes>
  );
}
