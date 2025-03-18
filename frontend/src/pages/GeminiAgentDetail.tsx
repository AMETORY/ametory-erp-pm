import { useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { useParams } from "react-router-dom";
import {
  generateContent,
  getGeminiAgentDetail,
  updateGeminiAgent,
} from "../services/api/geminiApi";
import {
  Button,
  Label,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";
import Select, { InputActionMeta } from "react-select";
import { GeminiAgent } from "../models/gemini";
import toast from "react-hot-toast";
import { Mention, MentionsInput } from "react-mentions";
import { RiAttachment2 } from "react-icons/ri";
interface GeminiAgentDetailProps {}

const GeminiAgentDetail: FC<GeminiAgentDetailProps> = ({}) => {
  const { agentId } = useParams();
  const [agent, setAgent] = useState<GeminiAgent>();
  const [emojis, setEmojis] = useState([]);
  const [content, setContent] = useState("");
  const [openAttachment, setOpenAttachment] = useState(false);

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
        .then((resp: any) => setAgent(resp.data))
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
            <div className="mb-2 block">
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
            </div>
          </form>
          <div className="flex justify-end gap-2 w-full">
            <Button
              onClick={() => {
                updateGeminiAgent(agentId!, agent)
                  .then(() => {
                    getGeminiAgentDetail(agentId!)
                      .then((resp: any) => setAgent(resp.data))
                      .catch(toast.error);
                  })
                  .catch(toast.error);
              }}
            >
              Create
            </Button>
          </div>
        </div>
        <div className="w-[calc(100%-300px)] border-l relative bg-gray-50">
          <div
            id="channel-messages"
            className="messages h-[calc(100vh-120px)] overflow-y-auto p-4 bg-red-50 "
          ></div>
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
                    await generateContent(content, agentId!).then(
                      (resp: any) => {
                        toast.success(resp.data.response);
                      }
                    );
                  } catch (error) {
                    toast.error(`${error}`);
                  } finally {
                    setTimeout(() => {
                      setContent("");
                    }, 300);
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
    </AdminLayout>
  );
};
export default GeminiAgentDetail;
