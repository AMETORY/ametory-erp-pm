import { useEffect, useState, type FC } from "react";
import { asyncStorage } from "../utils/async_storage";
import {
  authorizeConnection,
  updateConnection,
} from "../services/api/connectionApi";
import { TbBrandShopee } from "react-icons/tb";

interface ShopeeAuthProps {}

const ShopeeAuth: FC<ShopeeAuthProps> = ({}) => {
  const urlParams = new URLSearchParams(window.location.search);
  const shopeeCode = urlParams.get("code");
  const shopID = urlParams.get("shop_id");
  const [mounted, setMounted] = useState(false);
  const [shops, setShops] = useState<any[]>([]);
  const [succeed, setSucceed] = useState(false);

  useEffect(() => {
    setMounted(true);

    return () => {};
  }, []);

  const [connectionID, setConnectionID] = useState("");

  useEffect(() => {
    if (!mounted) return;
    asyncStorage.getItem("shopee-connection-id").then((id) => {
      setConnectionID(id);
    });
  }, [shopeeCode, mounted]);

  useEffect(() => {
    
  }, [shopeeCode, connectionID]);

  const authorize = async () => {
    if (shopeeCode && shopID) {
      authorizeConnection(connectionID, {
        shop_id: shopID,
        shopee_code: shopeeCode,
        type: "shopee",
      })
        .then((resp: any) => {
          // console.log("Authorized", resp.data);
          setSucceed(true);
          window.location.href = "/connection";
        })
        .catch(console.error);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <TbBrandShopee size={72} />
      <p className="text-center text-2xl">
        Please wait for the next step in authorization...
      </p>
      {/* <p>Connection ID : {connectionID}</p>
      <p>Code : {shopeeCode}</p> */}
      <p>SHOP ID : {shopID}</p>

      {shopeeCode && shopID && (
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
export default ShopeeAuth;
