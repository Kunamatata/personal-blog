import { useState } from "react";
import { useQuery } from "react-query";
import { Route, Routes } from "react-router-dom";
import { Navbar } from "./components/Navbar/Navbar";
import { PostPage } from "./components/PostPage/PostPage";
import { Posts } from "./components/Posts/Posts";

export const App = () => {
  return (
    <>
      <Navbar />
      <div className="p-6 md:max-w-screen-md md:p-0 m-auto transition-all duration-500">
        <Routes>
          <Route path="/" element={<Posts />} />
          <Route path="/posts/:slug" element={<PostPage />} />
        </Routes>
      </div>
    </>
  );
};
