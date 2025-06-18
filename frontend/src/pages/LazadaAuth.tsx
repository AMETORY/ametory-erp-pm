import type { FC } from "react";

interface LazadaAuthProps {}

const LazadaAuth: FC<LazadaAuthProps> = ({}) => {
  const urlParams = new URLSearchParams(window.location.search);
  const code = urlParams.get("code");
  const shopID = urlParams.get("shop_id");

  const authorize = async () => {
    if (code && shopID) {
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <img src="/lazada.png" alt="" className="h-16" />
      <p className="text-center text-2xl">
        Please wait for the next step in authorization...
      </p>


      {(code && shopID) && (
        <button
          className="mt-4 px-4 py-2 bg-blue-500 hover:bg-blue-700 text-white font-bold rounded w-64"
          onClick={authorize}
        >
          Authorize Now
        </button>
      )}
    </div>
  );
};
export default LazadaAuth;
