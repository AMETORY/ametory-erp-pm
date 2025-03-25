import { useContext, useEffect, useState, type FC } from "react";
import { ConnectionModel } from "../models/connection";
import {
  Badge,
  Button,
  ButtonGroup,
  Label,
  Spinner,
  ToggleSwitch,
} from "flowbite-react";
import QRCode from "react-qr-code";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { useParams } from "react-router-dom";
import { LoadingContext } from "../contexts/LoadingContext";
import {
  connectDevice,
  getConnection,
  getQr,
  syncContactConnection,
} from "../services/api/connectionApi";
import toast from "react-hot-toast";
import { getGeminiAgents } from "../services/api/geminiApi";
import { GeminiAgent } from "../models/gemini";
import Select, { InputActionMeta } from "react-select";
import { BsCheck2, BsCheck2Circle } from "react-icons/bs";
import { FaCircleXmark } from "react-icons/fa6";
import { ContactModel } from "../models/contact";

interface ConnectionDrawerProps {
  connection: ConnectionModel;
  onUpdate: (connection: ConnectionModel) => void;
  onFinish: () => void;
  onSave: () => void;
  
}

const ConnectionDrawer: FC<ConnectionDrawerProps> = ({
  connection,
  onUpdate,
  onFinish,
  onSave,
  
}) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const [mounted, setMounted] = useState(false);
  const { loading, setLoading } = useContext(LoadingContext);
  const [qrLoading, setQrLoading] = useState(false);
  const [qrStr, setQrStr] = useState("");
  const [geminiAgents, setGeminiAgents] = useState<GeminiAgent[]>([]);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted && connection.id) {
      //   getConnection(connection.id).then((res: any) => {
      //     setConnection(res.data);
      //   });
      getGeminiAgents({ page: 1, size: 20 }).then((res: any) => {
        setGeminiAgents(res.data.items);
      });
    }
  }, [mounted, connection.id]);

  return (
    <div className=" flex flex-col mt-16">
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
      {connection?.status == "ACTIVE" && connection.type == "whatsapp" && (
        <div className="mt-4">
          <Label className="font-bold">Auto Pilot</Label>
          <ToggleSwitch
            checked={connection?.is_auto_pilot ?? false}
            onChange={(e) => {
              onUpdate({
                ...connection!,
                is_auto_pilot: e,
              });
            }}
          />
        </div>
      )}
      {connection?.status == "ACTIVE" && connection.type == "whatsapp" && (
        <div className="mt-4">
          <Label className="font-bold">Session Auth</Label>
          <ToggleSwitch
            checked={connection?.session_auth ?? false}
            onChange={(e) => {
              onUpdate({
                ...connection!,
                session_auth: e,
              });
            }}
          />
        </div>
      )}
      {connection?.status == "ACTIVE" && connection.type == "whatsapp" && (
        <div className="mt-4">
          <Label className="font-bold">Gemini Agent</Label>
          <Select
            options={geminiAgents}
            value={connection?.gemini_agent}
            formatOptionLabel={(option) => (
              <div>
                <p>{option.name}</p>
              </div>
            )}
            styles={{
              container: (provided) => ({
                ...provided,
                width: "100%",
                borderRadius: "5px",
              }),
            }}
            onChange={(e) => {
              onUpdate({
                ...connection!,
                gemini_agent_id: e?.id,
                gemini_agent: e!,
              });
            }}
          />
        </div>
      )}
      {connection?.status == "ACTIVE" && connection.type == "whatsapp" && (
        <div className="mt-4">
          <Label className="font-bold">Device Connection</Label>
          <p>
            {connection?.connected ? (
              <div className="flex gap-2 items-center">
                <BsCheck2Circle className="text-green-500" /> Connected
              </div>
            ) : (
              <div className="flex gap-2 items-center">
                <FaCircleXmark className="text-red-500" /> Disconnected
              </div>
            )}
          </p>
        </div>
      )}
      {connection?.status == "ACTIVE" && connection.type == "whatsapp" && (
        <div className="mt-4"></div>
      )}
      {connection?.status == "ACTIVE" && connection.type == "whatsapp" && (
        <div className="mt-4">
          <ButtonGroup>
            <Button
              className="mt-4"
              color="warning"
              onClick={async () => {
                syncContactConnection(connection.id!)
                  .then((res: any) => {
                    toast.success("Sync Success");
                  })
                  .catch((err: any) => {
                    toast.error("Sync Failed");
                  });
              }}
            >
              Sync Contact
            </Button>
            <Button
              className="mt-4"
              onClick={async () => {
                onSave();
              }}
            >
              SAVE
            </Button>
          </ButtonGroup>
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

          {qrStr != "" && !qrStr.includes("redis") && <QRCode value={qrStr} />}
          {qrLoading && <Spinner aria-label="Default status example" />}
          <Button
            className="mt-4"
            onClick={async () => {
              var refreshIntervalId: any;
              setQrLoading(true);
              connectDevice(connection!.id!)
                .then((res: any) => {
                  onFinish();
                })
                .catch((err) => {
                  toast.error(`${err}`);
                })

                .finally(() => {
                  setQrLoading(false);
                  clearInterval(refreshIntervalId);
                });

              setTimeout(() => {
                getQr(connection!.id!, connection!.session_name!)
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
              refreshIntervalId = setInterval(async () => {
                setQrLoading(true);
                getQr(connection!.id!, connection!.session_name!)
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
  );
};
export default ConnectionDrawer;
