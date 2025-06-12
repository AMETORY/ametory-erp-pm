import { useEffect, useState, type FC } from "react";
import { PiTiktokLogoDuotone } from "react-icons/pi";
import { asyncStorage } from "../utils/async_storage";
import {
  authorizeConnection,
  updateConnection,
} from "../services/api/connectionApi";

interface TiktokAuthProps {}

const TiktokAuth: FC<TiktokAuthProps> = ({}) => {
  const urlParams = new URLSearchParams(window.location.search);
  const tiktokCode = urlParams.get("code");
  const tiktokState = urlParams.get("state");
  const [mounted, setMounted] = useState(false);
  const [shops, setShops] = useState<any[]>([]);

  useEffect(() => {
    setMounted(true);

    return () => {};
  }, []);

  const [connectionID, setConnectionID] = useState("");

  useEffect(() => {
    if (!mounted) return;
    asyncStorage.getItem("tiktok-connection-id").then((id) => {
      setConnectionID(id);
      authorizeConnection(id, {
        tiktok_code: tiktokCode,
        type: "tiktok",
      })
        .then((resp: any) => {
          // console.log("Authorized", resp.data);
          // window.location.href = "/connection";
          setShops(resp.data);
        })
        .catch(console.error);
    });
  }, [tiktokCode, mounted]);

  useEffect(() => {
    if (tiktokCode && tiktokState) {
      console.log(tiktokCode, connectionID);
    }
  }, [tiktokCode, connectionID]);

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <PiTiktokLogoDuotone size={72} />
      {shops.length > 0 && (
        <p className="text-center text-2xl">Please select Active Shop</p>
      )}
      {shops.map((shop) => {
        return (
          <div
            className="max-w-[300px] hover:bg-gray-50 rounded-lg p-4 mb-4 border border-gray-200 cursor-pointer flex flex-col"
            key={shop.id}
            onClick={() => {
              updateConnection(connectionID, {
                session_name: shop.name,
                username: shop.id,
                password: shop.cipher,
              }).then(() => {
                asyncStorage.removeItem("tiktok-connection-id").then(() => {
                  window.location.href = "/connection";
                });
              });
            }}
          >
            <h3>{shop.name}</h3>
            <small>Reg: {shop.region}</small>
            <small>Type: {shop.seller_type}</small>
          </div>
        );
      })}
      {shops.length == 0 && (
        <p className="text-center text-2xl">
          Please wait for the next step in authorization...
        </p>
      )}
    </div>
  );
};
export default TiktokAuth;
