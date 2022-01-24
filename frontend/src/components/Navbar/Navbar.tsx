import { NavLink } from "react-router-dom";

export const Navbar = () => {
  return (
    <div className="flex gap-8 justify-center md:justify-end p-8 md:mr-20">
      <div className="text-2xl">
        <NavLink
          to="/"
          className={(link) =>
            link.isActive
              ? "text-active-link  border-b-2 border-active-link"
              : "text-white"
          }
        >
          home
        </NavLink>
      </div>
      <div className=" text-2xl">
        <NavLink
          to="/about"
          className={(link) =>
            link.isActive
              ? "text-active-link border-b-2 border-active-link"
              : "text-white"
          }
        >
          about
        </NavLink>
      </div>
      <div className=" text-2xl">
        <NavLink
          to="/projects"
          className={(link) =>
            link.isActive
              ? "text-active-link  border-b-2 border-active-link"
              : "text-white"
          }
        >
          projects
        </NavLink>
      </div>
    </div>
  );
};
