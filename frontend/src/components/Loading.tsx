import type { FC } from "react";
import { AiOutlineLoading3Quarters } from "react-icons/ai";

interface LoadingProps {}

const Loading: FC<LoadingProps> = ({}) => {
  return (
    <div className="fixed inset-0 flex items-center justify-center bg-white bg-opacity-80 loading z-50">
      <div className="animate-spin h-5 w-5 border-b-2 border-gray-900 rounded-full">
        <AiOutlineLoading3Quarters />
      </div>
    </div>
  );
};
export default Loading;
