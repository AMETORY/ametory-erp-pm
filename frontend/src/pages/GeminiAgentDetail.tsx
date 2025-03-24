import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { useParams } from "react-router-dom";
import {
  deleteGeminiAgentHistory,
  generateContent,
  getGeminiAgentDetail,
  getGeminiAgentHistories,
  toggleGeminiAgentHistoryModel,
  updateGeminiAgent,
  updateGeminiAgentHistory,
} from "../services/api/geminiApi";
import {
  Button,
  Label,
  Modal,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";
import Select, { InputActionMeta } from "react-select";
import { GeminiAgent, GeminiAgentHistory } from "../models/gemini";
import toast from "react-hot-toast";
import { Mention, MentionsInput } from "react-mentions";
import { RiAttachment2 } from "react-icons/ri";
import { error } from "console";
import {
  allExpanded,
  darkStyles,
  defaultStyles,
  JsonView,
} from "react-json-view-lite";
import "react-json-view-lite/dist/index.css";
import { LoadingContext } from "../contexts/LoadingContext";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { BsTrash } from "react-icons/bs";
import { JsonEditor } from "json-edit-react";
interface GeminiAgentDetailProps {}

const GeminiAgentDetail: FC<GeminiAgentDetailProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const { agentId } = useParams();
  const [agent, setAgent] = useState<GeminiAgent>();
  const [emojis, setEmojis] = useState([]);
  const [content, setContent] = useState("");
  const [openAttachment, setOpenAttachment] = useState(false);
  const [histories, setHistories] = useState<GeminiAgentHistory[]>([]);
  const [showHtml, setShowHtml] = useState(false);
  const [activeHistory, setActiveHistory] = useState<GeminiAgentHistory>();

  const emojiStyle = {
    control: {
      fontSize: 16,
      lineHeight: 1.2,
      minHeight: 30,
      maxHeight: 80,
    },

    highlighter: {
      padding: 9,
      border: "1px solid transparent",
    },

    input: {
      fontSize: 16,
      lineHeight: 1.2,
      padding: 9,
      border: "1px solid silver",
      borderRadius: 10,
    },

    suggestions: {
      list: {
        backgroundColor: "white",
        border: "1px solid rgba(0,0,0,0.15)",
        fontSize: 16,
      },

      item: {
        padding: "5px 15px",
        borderBottom: "1px solid rgba(0,0,0,0.15)",

        "&focused": {
          backgroundColor: "#cee4e5",
        },
      },
    },
  };

  useEffect(() => {
    if (agentId) {
      getGeminiAgentDetail(agentId)
        .then((resp: any) => {
          setAgent(resp.data);
          getDetail();
        })
        .catch(toast.error);
    }
  }, [agentId]);

  useEffect(() => {
    fetch(
      "https://gist.githubusercontent.com/oliveratgithub/0bf11a9aff0d6da7b46f1490f86a71eb/raw/d8e4b78cfe66862cf3809443c1dba017f37b61db/emojis.json"
    )
      .then((response) => {
        return response.json();
      })
      .then((jsonData) => {
        setEmojis(jsonData.emojis);
      });
  }, []);

  useEffect(() => {
    if (agent) {
    }
  }, [agent]);

  const getDetail = () => {
    getGeminiAgentHistories(agentId!)
      .then((resp: any) => {
        setHistories(resp.data);
        setTimeout(scrollToBottom, 300);
      })
      .catch(toast.error);
  };

  const neverMatchingRegex = /($a)/;
  const queryEmojis = (query: any, callback: (emojis: any) => void) => {
    if (query.length === 0) return;

    const matches = emojis
      .filter((emoji: any) => {
        return emoji.name.indexOf(query.toLowerCase()) > -1;
      })
      .slice(0, 10);
    return matches.map(({ emoji }) => ({ id: emoji }));
  };

  const geminiModels = [
    { label: "Gemini 1.0", value: "gemini-1.0" },
    { label: "Gemini 1.5", value: "gemini-1.5" },
    { label: "Gemini 2.0", value: "gemini-2.0" },
    { label: "Gemini 2.0 Flash", value: "gemini-2.0-flash" },
    { label: "Gemini 2.0 Flash Exp", value: "gemini-2.0-flash-exp" },
  ];

  const scrollToBottom = () => {
    const element = document.getElementById("messages");
    if (element) {
      element.scrollTo({
        top: element.scrollHeight,
        behavior: "smooth",
      });
    }
  };

  return (
    <AdminLayout>
      <div className="flex flex-row w-full h-full flex-1 gap-2">
        <div className="w-[300px] h-full p-4 space-y-4 flex flex-col overflow-y-auto">
          <div className="flex justify-between items-center">
            <h3 className="text-2xl font-bold">Agent Detail</h3>
            <div className="flex gap-2 items-center"></div>
          </div>
          <form>
            <div className="mb-2 block">
              <Label htmlFor="name" value="Name" />
              <TextInput
                id="name"
                type="text"
                value={agent?.name ?? ""}
                onChange={(e) =>
                  setAgent((prev) => ({
                    ...prev,
                    name: e.target.value,
                  }))
                }
                placeholder="Name"
                required={true}
              />
            </div>
            <div className="mb-2 block">
              <Label htmlFor="system-instruction" value="System Instructon" />
              <Textarea
                id="system-instruction"
                value={agent?.system_instruction ?? ""}
                onChange={(e) =>
                  setAgent((prev) => ({
                    ...prev,
                    system_instruction: e.target.value,
                  }))
                }
                placeholder="System Instructon"
                rows={5}
                required={true}
              />
            </div>
            <div className="mb-2 block">
              <Label htmlFor="api-key" value="API Key" />
              <TextInput
                id="api-key"
                type="text"
                value={agent?.api_key ?? ""}
                onChange={(e) =>
                  setAgent((prev) => ({
                    ...prev,
                    api_key: e.target.value,
                  }))
                }
                placeholder="API Key"
                required={true}
              />
            </div>
            <div className="mb-2 block">
              <Label htmlFor="active" value="Active" />
              <ToggleSwitch
                id="active"
                checked={agent?.active ?? false}
                onChange={(checked) =>
                  setAgent((prev) => ({
                    ...prev,
                    active: checked,
                  }))
                }
              />
            </div>

            <div className="mb-2 block">
              <Label htmlFor="model" value="Model" />
              <Select
                id="model"
                value={geminiModels.find((v) => v.value == agent?.model)}
                onChange={(option) =>
                  setAgent((prev) => ({
                    ...prev,
                    model: option?.value,
                  }))
                }
                options={geminiModels}
                required={true}
              />
            </div>
            <div className="mb-2 block">
              <Label htmlFor="set-temperature" value="Set Temperature" />
              <TextInput
                id="set-temperature"
                type="number"
                value={agent?.set_temperature ?? 1}
                onChange={(e) =>
                  setAgent((prev) => ({
                    ...prev,
                    set_temperature: Number(e.target.value),
                  }))
                }
                placeholder="Set Temperature"
                required={true}
              />
            </div>
            <div className="mb-2 block">
              <Label htmlFor="set-top-k" value="Set Top K" />
              <TextInput
                id="set-top-k"
                type="number"
                value={agent?.set_top_k ?? 40}
                onChange={(e) =>
                  setAgent((prev) => ({
                    ...prev,
                    set_top_k: Number(e.target.value),
                  }))
                }
                placeholder="Set Top K"
                required={true}
              />
            </div>
            <div className="mb-2 block">
              <Label htmlFor="set-top-p" value="Set Top P" />
              <TextInput
                id="set-top-p"
                type="number"
                value={agent?.set_top_p ?? 0.95}
                onChange={(e) =>
                  setAgent((prev) => ({
                    ...prev,
                    set_top_p: Number(e.target.value),
                  }))
                }
                placeholder="Set Top P"
                required={true}
              />
            </div>
            <div className="mb-2 block">
              <Label
                htmlFor="set-max-output-tokens"
                value="Set Max Output Tokens"
              />
              <TextInput
                id="set-max-output-tokens"
                type="number"
                value={agent?.set_max_output_tokens ?? 8192}
                onChange={(e) =>
                  setAgent((prev) => ({
                    ...prev,
                    set_max_output_tokens: Number(e.target.value),
                  }))
                }
                placeholder="Set Max Output Tokens"
                required={true}
              />
            </div>
            {/* <div className="mb-2 block">
              <Label htmlFor="response-mimetype" value="Response Mimetype" />
              <TextInput
                id="response-mimetype"
                type="text"
                value={agent?.response_mimetype ?? "application/json"}
                onChange={(e) =>
                  setAgent((prev) => ({
                    ...prev,
                    response_mimetype: e.target.value,
                  }))
                }
                placeholder="Response Mimetype"
                required={true}
              />
            </div> */}
            <div className="mb-2 block">
              <Label htmlFor="html" value="View HTML" />
              <ToggleSwitch
                id="html"
                checked={showHtml}
                onChange={(checked) => setShowHtml(checked)}
              />
            </div>
          </form>
          <div className="flex justify-end gap-2 w-full">
            <Button
              onClick={() => {
                updateGeminiAgent(agentId!, agent)
                  .then(() => {
                    toast.success("Agent updated");
                    getGeminiAgentDetail(agentId!)
                      .then((resp: any) => setAgent(resp.data))
                      .catch(toast.error);
                  })
                  .catch(toast.error);
              }}
            >
              Save
            </Button>
          </div>
        </div>
        <div className="w-[calc(100%-300px)] border-l relative bg-gray-50">
          <div
            id="messages"
            className="messages h-[calc(100vh-120px)] overflow-y-auto p-4 "
          >
            {histories.map((e) => (
              <div
                key={e.id}
                className="space-y-8 group/item hover:bg-yellow-50 p-2"
              >
                <div className="flex flex-row justify-between  items-center">
                  <div>
                    <small className="ml-2">User</small>
                    <div className="bg-white rounded-lg  p-4 max-w-[600px]">
                      {e.input}
                    </div>
                  </div>
                  <div className="flex flex-row gap-2 group/edit invisible group-hover/item:visible">
                    <ToggleSwitch
                      sizing="sm"
                      label={(e.is_model && "Active") || "Inactive"}
                      checked={e.is_model ?? false}
                      onChange={(checked) => {
                        toggleGeminiAgentHistoryModel(agentId!, e.id!);
                        setHistories(
                          histories.map((item) => {
                            if (item.id === e.id) {
                              return { ...item, is_model: checked };
                            }
                            return item;
                          })
                        );
                      }}
                    />

                    <BsTrash
                      className="text-red-400 hover:text-red-600 cursor-pointer "
                      onClick={() => {
                        if (
                          window.confirm(
                            "Are you sure you want to delete this history?"
                          )
                        ) {
                          deleteGeminiAgentHistory(agentId!, e.id!).then(() => {
                            setHistories([
                              ...histories.filter((item) => item.id !== e.id),
                            ]);
                          });
                        }
                      }}
                    />
                  </div>
                </div>
                <div className="flex flex-row justify-end">
                  <div className="flex flex-col  items-end">
                    <div className="flex gap-2">
                      <small className="mr-2">Model</small>
                      <small
                        className="mr-2 cursor-pointer hover:underline"
                        onClick={() => {
                          setActiveHistory(e);
                        }}
                      >
                        Edit
                      </small>
                    </div>
                    <div className="bg-white rounded-lg  p-4 max-w-[600px]">
                      <div className=" bg-white rounded-lg p-4 max-w-[600px] json-container">
                        {showHtml ? (
                          <Markdown remarkPlugins={[remarkGfm]}>
                            {JSON.parse(e.output).response ?? ""}
                          </Markdown>
                        ) : (
                          <JsonView
                            data={JSON.parse(e.output!)}
                            shouldExpandNode={allExpanded}
                            style={{
                              ...defaultStyles,
                            }}
                          />
                        )}
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            ))}
          </div>
          <div className="shoutbox border-t pt-2 min-h-[20px] max-h[60px] px-2  flex justify-between items-center gap-2 bg-white">
            <MentionsInput
              value={content}
              onChange={(val: any) => {
                setContent(val.target.value);
              }}
              style={emojiStyle}
              placeholder={
                "Press ':' for emojis, mention people using '@' and shift+enter to send"
              }
              className="w-full"
              autoFocus
              onKeyDown={async (val: any) => {
                if (val.key === "Enter" && val.shiftKey) {
                  try {
                    // await createMessage(channelId!, {
                    //   message: content,
                    //   files: files,
                    // });
                    // setOpenAttachment(false);
                    // setFiles([]);
                    setLoading(true);
                    setTimeout(() => {
                      setContent("");
                    }, 300);
                    await generateContent(content, agentId!, false, false).then(
                      (resp: any) => {
                        // toast.success(resp.data.response);
                        getGeminiAgentHistories(agentId!)
                          .then((resp: any) => setHistories(resp.data))
                          .catch(toast.error)
                          .finally(() => setLoading(false));
                      }
                    );
                  } catch (error) {
                    toast.error(`${error}`);
                  } finally {
                    setTimeout(scrollToBottom, 300);
                  }

                  return;
                }
              }}
            >
              {/* <Mention
                trigger="@"
                data={(channel?.participant_members ?? []).map((member) => ({
                  id: member.id!,
                  display: member.user?.full_name!,
                }))}
                style={{
                  backgroundColor: "#cee4e5",
                }}
                appendSpaceOnAdd
              /> */}
              <Mention
                trigger=":"
                markup="__id__"
                regex={neverMatchingRegex}
                data={queryEmojis}
              />
            </MentionsInput>
            <Button color="gray" onClick={() => setOpenAttachment(true)}>
              <RiAttachment2 />
            </Button>
          </div>
        </div>
      </div>
      {activeHistory && (
        <Modal
          show={activeHistory !== undefined}
          onClose={() => setActiveHistory(undefined)}
        >
          <Modal.Header>Edit History</Modal.Header>
          <Modal.Body>
            <div className="flex flex-col space-y-4">
              <div>
                <Label>Input</Label>
                <TextInput
                  value={activeHistory.input}
                  onChange={(e) => {
                    setActiveHistory({
                      ...activeHistory,
                      input: e.target.value,
                    });
                  }}
                />
              </div>
              <div>
                <Label>Input</Label>
                <JsonEditor
                  data={JSON.parse(activeHistory!.output!)}
                  setData={(val) => {
                    setActiveHistory({
                      ...activeHistory!,
                      output: JSON.stringify(val),
                    });
                  }}
                />
              </div>
            </div>
          </Modal.Body>
          <Modal.Footer>
            <div className="flex w-full justify-end">
              <Button
                onClick={() => {
                  setLoading(true);
                  updateGeminiAgentHistory(
                    agentId!,
                    activeHistory!.id!,
                    activeHistory
                  )
                    .catch(toast.error)
                    .finally(() => {
                      setLoading(false);
                      setHistories([
                        ...histories.map((item) => {
                          if (item.id == activeHistory.id) {
                            return activeHistory;
                          }
                          return item;
                        }),
                      ]);
                      setActiveHistory(undefined);
                    });
                }}
              >
                Save
              </Button>
            </div>
          </Modal.Footer>
        </Modal>
      )}
    </AdminLayout>
  );
};
export default GeminiAgentDetail;
