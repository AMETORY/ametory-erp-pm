import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { useNavigate, useParams } from "react-router-dom";
import { WebsocketContext } from "../contexts/WebsocketContext";
import {
  connectDevice,
  getConnection,
  getQr,
} from "../services/api/connectionApi";
import { ConnectionModel } from "../models/connection";
import { Badge, Button, Label, Spinner } from "flowbite-react";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";
import QRCode from "react-qr-code";

interface ConnectionDetailProps {}

const ConnectionDetail: FC<ConnectionDetailProps> = ({}) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const [mounted, setMounted] = useState(false);
  const { connectionId } = useParams();
  const [connection, setConnection] = useState<ConnectionModel>();
  const { loading, setLoading } = useContext(LoadingContext);
  const [qrLoading, setQrLoading] = useState(false);
  const [qrStr, setQrStr] = useState("");

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted && connectionId) {
      getConnection(connectionId).then((res: any) => {
        setConnection(res.data);
      });
    }
  }, [mounted, connectionId]);
  return (
    <AdminLayout>
      <div className="p-8 flex flex-col">
        <h3 className="text-2xl font-bold">{connection?.name}</h3>
        <p className="text-lg">{connection?.description}</p>
        <div className="flex w-fit">
          {connection?.status == "PENDING" && (
            <Badge color="warning">{connection?.status}</Badge>
          )}
          {connection?.status == "ACTIVE" && (
            <Badge color="success">{connection?.status}</Badge>
          )}
        </div>
        {connection?.status == "ACTIVE" && connection.type == "whatsapp" && (
          <div className="mt-4">
            <Label className="font-bold">Phone Number</Label>
            <p>{connection?.session_name}</p>
          </div>
        )}
        {connection?.status == "PENDING" && connection.type == "whatsapp" && (
          <div className="mt-4 p-4 border rounded-lg">
            <h1 className="text-2xl font-bold">Connect to WhatsApp</h1>
            <table style={{}}>
              <tbody>
                <tr>
                  <td className="less-pad">
                    1.&nbsp;&nbsp;Open{" "}
                    <span
                      className="grey lighten-3 black-text"
                      style={{
                        padding: "2px 5px",
                        margin: "0 3px",
                        borderRadius: 5,
                      }}
                    >
                      WhatsApp
                    </span>{" "}
                    on your phone
                  </td>
                </tr>
                <tr>
                  <td className="less-pad">
                    2.&nbsp;&nbsp;Click{" "}
                    <span
                      className="grey lighten-3 black-text"
                      style={{
                        padding: "2px 5px",
                        margin: "0 3px",
                        borderRadius: 5,
                      }}
                    >
                      3-dots
                    </span>{" "}
                    menu on the top right corner
                  </td>
                </tr>
                <tr>
                  <td className="less-pad">
                    3.&nbsp;&nbsp;Tap on{" "}
                    <span
                      className="grey lighten-3 black-text"
                      style={{
                        padding: "2px 5px",
                        margin: "0 3px",
                        borderRadius: 5,
                      }}
                    >
                      "Linked Devices"
                    </span>
                  </td>
                </tr>
                <tr>
                  <td className="less-pad">
                    4.&nbsp;&nbsp;After that,{" "}
                    <span
                      className="grey lighten-3 black-text"
                      style={{
                        padding: "2px 5px",
                        margin: "0 3px",
                        borderRadius: 5,
                      }}
                    >
                      Scan QR Code
                    </span>{" "}
                    below
                  </td>
                </tr>
              </tbody>
            </table>

            {qrStr != "" && !qrStr.includes("redis") && (
              <QRCode value={qrStr} />
            )}
            {qrLoading && <Spinner aria-label="Default status example" />}
            <Button
              className="mt-4"
              onClick={async () => {
                setQrLoading(true);
                connectDevice(connectionId!)
                  .catch((err) => {
                    toast.error(`${err}`);
                  })
                  .then((res: any) => {
                    window.location.reload();
                  })
                  .finally(() => {
                    setQrLoading(false);
                  });

                setTimeout(() => {
                  getQr(connectionId!, connection!.session_name!)
                    .then((res: any) => {
                      setQrStr("");
                      setQrStr((val) => res.data);
                    })
                    .catch((err) => {
                      toast.error(`${err}`);
                      setQrStr("");
                    })
                    .finally(() => {
                      setQrLoading(false);
                    });
                }, 3000);
                setInterval(async () => {
                  setQrLoading(true);
                  getQr(connectionId!, connection!.session_name!)
                    .then((res: any) => {
                      setQrStr("");
                      setQrStr((val) => res.data);
                    })
                    .catch((err) => {
                      toast.error(`${err}`);
                    })
                    .finally(() => {
                      setQrLoading(false);
                    });
                }, 10000);
              }}
            >
              Connect Device
            </Button>
          </div>
        )}
      </div>
    </AdminLayout>
  );
};
export default ConnectionDetail;
