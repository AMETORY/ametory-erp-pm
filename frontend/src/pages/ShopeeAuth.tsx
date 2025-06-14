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
    if (shopeeCode && shopID) {

       authorizeConnection("83f47480-ea4e-4112-8171-ae73e047cb91", {
        shop_id: shopID,
        shopee_code: shopeeCode,
        type: "shopee",
      })
        .then((resp: any) => {
          // console.log("Authorized", resp.data);
          window.location.href = "/connection";
        })
        .catch(console.error);
    }
  }, [shopeeCode, connectionID]);

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <TbBrandShopee size={72} />
      <p className="text-center text-2xl">
        Please wait for the next step in authorization...
      </p>
      <p>Connection ID : {connectionID}</p>
      <p>Code : {shopeeCode}</p>
      <p>SHOP ID : {shopID}</p>
    </div>
  );
};
export default ShopeeAuth;
