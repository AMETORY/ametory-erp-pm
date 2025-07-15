import { useContext, useEffect, useState, type FC } from "react";
import { ConnectionModel } from "../models/connection";
import {
  Badge,
  Button,
  ButtonGroup,
  Label,
  Spinner,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";
import QRCode from "react-qr-code";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { useNavigate, useParams } from "react-router-dom";
import { LoadingContext } from "../contexts/LoadingContext";
import {
  connectDevice,
  getConnection,
  getLazadaAuthURL,
  getQr,
  getShopeeAuthURL,
  syncContactConnection,
} from "../services/api/connectionApi";
import toast from "react-hot-toast";
import { getGeminiAgents } from "../services/api/geminiApi";
import { GeminiAgent } from "../models/gemini";
import Select, { InputActionMeta } from "react-select";
import {
  BsCheck2,
  BsCheck2Circle,
  BsInstagram,
  BsTelegram,
  BsTiktok,
  BsWhatsapp,
} from "react-icons/bs";
import { FaCircleXmark } from "react-icons/fa6";
import { ContactModel } from "../models/contact";
import { getProjects } from "../services/api/projectApi";
import { ProjectModel } from "../models/project";
import { ColumnModel } from "../models/column";
import TelegramIntegrationGuide from "./TelegramGuide";
import { LuInstagram } from "react-icons/lu";
import { asyncStorage } from "../utils/async_storage";

interface ConnectionDrawerProps {
  connection: ConnectionModel;
  onUpdate: (connection: ConnectionModel) => void;
  onFinish: () => void;
  onSave: () => void;
  onAuthorize: () => void;
}

const ConnectionDrawer: FC<ConnectionDrawerProps> = ({
  connection,
  onUpdate,
  onFinish,
  onSave,
  onAuthorize,
}) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const [mounted, setMounted] = useState(false);
  const { loading, setLoading } = useContext(LoadingContext);
  const [qrLoading, setQrLoading] = useState(false);
  const [qrStr, setQrStr] = useState("");
  const [geminiAgents, setGeminiAgents] = useState<GeminiAgent[]>([]);
  const [projects, setProjects] = useState<ProjectModel[]>([]);
  const [columns, setColumns] = useState<ColumnModel[]>([]);
  const nav = useNavigate();

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
      getAllProjects("");
    }
  }, [mounted, connection.id]);

  const getAllProjects = (s: string) => {
    getProjects({ page: 1, size: 10, search: s })
      .then((e: any) => setProjects(e.data.items))
      .catch(toast.error);
  };
  const getIcon = (type?: string) => {
    switch (type) {
      case "whatsapp":
        return <BsWhatsapp className="mr-2 w-4" />;
      case "tiktok":
        return <BsTiktok className="mr-2 w-4" />;
      case "telegram":
        return <BsTelegram className="mr-2 w-4" />;
      case "instagram":
        return <BsInstagram className="mr-2 w-4" />;
      case "gemini":
        return "📲";
    }
  };
  return (
    <div className=" flex flex-col mt-16">
      <h3 className="text-2xl font-bold flex items-center">
        {" "}
        {getIcon(connection?.type)} {connection?.name}
      </h3>
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
      {connection.type == "telegram" && (
        <>
          <div className="mt-4">
            <Label className="font-bold">BOT Name</Label>
            <p>{connection?.session_name}</p>
          </div>
          <div className="mt-4">
            <Label className="font-bold">BOT TOKEN</Label>
            <TextInput
              value={connection?.access_token}
              onChange={(e) => {
                onUpdate({
                  ...connection!,
                  access_token: e.target.value,
                });
              }}
            />
          </div>
          {connection?.status != "ACTIVE" && (
            <div>
              <Button
                className="mt-4"
                onClick={async () => {
                  onSave();
                }}
              >
                SAVE
              </Button>
            </div>
          )}
        </>
      )}
      {connection.type == "tiktok" && (
        <>
          <div className="flex gap-2 flex-row">
            <Button
              className="mt-4 bg-yellow-400"
              onClick={async () => {
                asyncStorage
                  .setItem("tiktok-connection-id", connection.id)
                  .then(() => {
                    window.location.href = `https://services.tiktokshop.com/open/authorize?service_id=${process.env.REACT_APP_TIKTOK_SERVICE_ID}`;
                  });
              }}
            >
              {connection?.status == "ACTIVE" ? "Re-Authorize" : "Authorize"}
            </Button>
            {/* <Button
                className="mt-4"
                onClick={async () => {
                  onSave();
                }}
              >
                SAVE
              </Button> */}
          </div>
        </>
      )}
      {connection.type == "tiktok" &&
        connection?.status == "ACTIVE" &&
        connection.session_name && (
          <>
            <div className="mt-4">
              <Label className="font-bold">Connected To</Label>
              <p>{connection?.session_name}</p>
            </div>
            <div className="mt-4">
              <Label className="font-bold">Shop ID</Label>
              <p>{connection?.username}</p>
            </div>
          </>
        )}

      {connection.type == "shopee" && (
        <>
          <div className="flex gap-2 flex-row">
            <Button
              className="mt-4 bg-yellow-400"
              onClick={async () => {
                asyncStorage
                  .setItem("shopee-connection-id", connection.id)
                  .then(() => {
                    getShopeeAuthURL().then((res: any) => {
                      window.location.href = res.data;
                    });
                  });
              }}
            >
              {connection?.status == "ACTIVE" ? "Re-Authorize" : "Authorize"}
            </Button>
            {/* <Button
                className="mt-4"
                onClick={async () => {
                  onSave();
                }}
              >
                SAVE
              </Button> */}
          </div>
        </>
      )}
     

      {connection.type == "shopee" &&
        connection?.status == "ACTIVE" &&
        connection.session_name && (
          <>
            <div className="mt-4">
              <Label className="font-bold">Connected To</Label>
              <p>{connection?.session_name}</p>
            </div>
           
          </>
        )}


      {connection.type == "instagram" && (
        <div className="mt-4">
          <Button
            className="mt-4"
            color="primary"
            style={{
              width: "100%",
              backgroundColor: "#e1306c",
              color: "white",
            }}
            onClick={() => {
              // Add your logic to connect to Facebook here
              window.open(
                `https://www.instagram.com/oauth/authorize?enable_fb_login=0&force_authentication=1&redirect_uri=https://app.senandika.web.id/api/v1/facebook/instagram/callback&state=connection_id-${connection.id}&client_id=1033935721571526&response_type=code&scope=instagram_business_basic%2Cinstagram_business_manage_messages%2Cinstagram_business_manage_comments%2Cinstagram_business_content_publish%2Cinstagram_business_manage_insights`
              );
            }}
          >
            <BsInstagram className="mr-2 w-4" /> Hubungkan ke Instagram
          </Button>
        </div>
      )}


       {connection.type == "lazada" && (
        <>
          <div className="flex gap-2 flex-row">
            <Button
              className="mt-4 bg-yellow-400"
              onClick={async () => {
                asyncStorage
                  .setItem("lazada-connection-id", connection.id)
                  .then(() => {
                    getLazadaAuthURL().then((res: any) => {
                      // console.log("Lazada Auth URL", res.data);
                      window.location.href = res.data;
                    });
                  });
              }}
            >
              {connection?.status == "ACTIVE" ? "Re-Authorize" : "Authorize"}
            </Button>
            {/* <Button
                className="mt-4"
                onClick={async () => {
                  onSave();
                }}
              >
                SAVE
              </Button> */}
          </div>
        </>
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
          <Label className="font-bold">Auto Response Time</Label>
          <div className="grid grid-cols-2 gap-2">
            <TextInput
              type="time"
              value={connection?.auto_response_start_time}
              onChange={(e) => {
                onUpdate({
                  ...connection!,
                  auto_response_start_time: e.target.value,
                });
              }}
            />
            <TextInput
              type="time"
              value={connection?.auto_response_end_time}
              onChange={(e) => {
                onUpdate({
                  ...connection!,
                  auto_response_end_time: e.target.value,
                });
              }}
            />
          </div>
        </div>
      )}

      {connection?.status == "ACTIVE" && connection.type == "whatsapp" && (
        <div className="mt-4">
          <Label className="font-bold">Auto Response Message</Label>
          <Textarea
            value={connection?.auto_response_message}
            onChange={(e) => {
              onUpdate({
                ...connection!,
                auto_response_message: e.target.value,
              });
            }}
            placeholder="Enter auto response message"
          />
          <small className="" style={{ lineHeight: "0.8" }}>
            if you don't set any gemini agent auto response message will be sent
            to user during the time frame.
          </small>
        </div>
      )}

      {connection?.status == "ACTIVE" && connection.type == "whatsapp" && (
        <div className="mt-4">
          <Label className="font-bold">Gemini Agent</Label>
          <Select
            isClearable
            options={geminiAgents.map((e) => ({ label: e.name, value: e.id }))}
            value={{
              label: connection?.gemini_agent?.name,
              value: connection?.gemini_agent?.id,
            }}
            styles={{
              container: (provided) => ({
                ...provided,
                width: "100%",
                borderRadius: "5px",
              }),
            }}
            onChange={(e) => {
              let selected = geminiAgents.find(
                (agent) => agent.id == e?.value!
              );
              onUpdate({
                ...connection!,
                gemini_agent_id: selected?.id,
                gemini_agent: selected!,
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
      {connection?.status == "ACTIVE" && (
        <div className="mt-4">
          <div className=" block">
            <Label className="font-bold" htmlFor="name">
              Project
            </Label>
          </div>
          <Select
            value={
              connection?.project
                ? {
                    label: connection?.project?.name,
                    value: connection?.project?.id,
                  }
                : null
            }
            options={[
              { label: "Select Project", value: "" },
              ...projects.map((e) => ({
                label: e.name,
                value: e.id,
              })),
            ]}
            onChange={(val) => {
              onUpdate({
                ...connection!,
                project_id: val?.value,
                new_session_column: undefined,
                new_session_column_id: undefined,
                project: {
                  ...connection!.project!,
                  id: val!.value,
                  name: val!.label,
                  columns:
                    projects.find((e) => e.id == val?.value)?.columns ?? [],
                },
              });
            }}
            onInputChange={(val) => {
              getAllProjects(val);
            }}
          />
        </div>
      )}
      {connection?.project && (
        <div className="mt-4">
          <div className=" block">
            <Label className="font-bold" htmlFor="name">
              New Session Column
            </Label>
          </div>
          <Select
            value={
              connection?.new_session_column
                ? {
                    label: connection?.new_session_column?.name,
                    value: connection?.new_session_column?.id,
                  }
                : null
            }
            options={[
              { label: "Select Column", value: "" },
              ...(connection.project?.columns ?? []).map((e) => ({
                label: e.name!,
                value: e.id!,
              })),
            ]}
            onChange={(val) => {
              onUpdate({
                ...connection!,
                new_session_column_id: val?.value!,
                new_session_column: {
                  ...connection!.new_session_column!,
                  id: val!.value,
                  name: val!.label,
                },
              });
            }}
          />
        </div>
      )}
      {connection?.project && (
        <div className="mt-4">
          <div className=" block">
            <Label className="font-bold" htmlFor="name">
              Idle Column
            </Label>
          </div>
          <Select
            value={
              connection?.idle_column
                ? {
                    label: connection?.idle_column?.name,
                    value: connection?.idle_column?.id,
                  }
                : null
            }
            options={[
              { label: "Select Column", value: "" },
              ...(connection.project?.columns ?? []).map((e) => ({
                label: e.name!,
                value: e.id!,
              })),
            ]}
            onChange={(val) => {
              onUpdate({
                ...connection!,
                idle_column_id: val?.value!,
                idle_column: {
                  ...connection!.idle_column!,
                  id: val!.value,
                  name: val!.label,
                },
              });
            }}
          />
        </div>
      )}
      {connection?.project && (
        <div className="mt-4">
          <div className=" block">
            <Label className="font-bold" htmlFor="name">
              Idle Time (in days)
            </Label>
          </div>
          <TextInput
            type="number"
            value={connection?.idle_duration}
            onChange={(e) => {
              onUpdate({
                ...connection!,
                idle_duration: Number(e.target.value),
              });
            }}
          />
        </div>
      )}
      {(connection?.status == "PENDING" || (!connection?.connected && connection?.status == "ACTIVE")) && connection.type == "whatsapp" && (
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
      
      {connection?.status == "PENDING" && connection.type == "telegram" && (
        <TelegramIntegrationGuide />
      )}
    </div>
  );
};
export default ConnectionDrawer;
