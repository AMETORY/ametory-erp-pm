import type { FC } from "react";

interface logoProps {}

const logo: FC<logoProps> = ({}) => {
  return (
    <a href="#" className="flex ms-2 md:me-24 gap-2">
      <img
        src="/android-chrome-512x512.png"
        className="h-8 me-3"
        alt="Ametory Logo"
      />
      <span className="self-center text-xl font-semibold sm:text-2xl whitespace-nowrap dark:text-white">
        {process.env.REACT_APP_SITE_TITLE}
      </span>
    </a>
  );
};
export default logo;
