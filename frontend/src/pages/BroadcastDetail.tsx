import { useContext, useEffect, useState, type FC } from "react";
import { useParams } from "react-router-dom";
import AdminLayout from "../components/layouts/admin";
import { getBroadcast } from "../services/api/broadcastApi";
import { BroadcastModel } from "../models/broadcast";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";
import { Badge, Label, Textarea } from "flowbite-react";
import { parseMentions } from "../utils/helper-ui";
import { Mention, MentionsInput } from "react-mentions";

interface BroadcastDetailProps {}
const neverMatchingRegex = /($a)/;
const BroadcastDetail: FC<BroadcastDetailProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [emojis, setEmojis] = useState([]);

  const { broadcastId } = useParams();
  const [mounted, setMounted] = useState(false);
  const [broadcast, setBroadcast] = useState<BroadcastModel>();
  const [isEditable, setisEditable] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

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
  const queryEmojis = (query: any, callback: (emojis: any) => void) => {
    if (query.length === 0) return;

    const matches = emojis
      .filter((emoji: any) => {
        return emoji.name.indexOf(query.toLowerCase()) > -1;
      })
      .slice(0, 10);
    return matches.map(({ emoji }) => ({ id: emoji }));
  };

  useEffect(() => {
    if (mounted && broadcastId) {
      setLoading(true);
      getBroadcast(broadcastId)
        .then((res: any) => {
          setBroadcast(res.data);
          setisEditable(res.data.status === "DRAFT");
        })
        .catch((error) => {
          toast.error(`${error}`);
        })
        .finally(() => {
          setLoading(false);
        });
    }
  }, [mounted, broadcastId]);
  return (
    <AdminLayout>
      <div className="p-8">
        <h1 className="text-2xl font-bold">Detail Broadcast</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4">
          {broadcast && (
            <div className="bg-white border rounded p-6 flex flex-col space-y-4">
              <div>
                <Label>Description</Label>
                {isEditable ? (
                  <Textarea
                    value={broadcast?.description ?? ""}
                    onChange={(val) => {
                      setBroadcast({
                        ...broadcast,
                        description: val.target.value,
                      });
                    }}
                  />
                ) : (
                  <p className="">{broadcast.description}</p>
                )}
              </div>
              <div>
                <Label>Message</Label>
                <p className="">
                  {isEditable ? (
                    <MentionsInput
                      value={broadcast?.message ?? ""}
                      onChange={(val) => {
                        setBroadcast({
                          ...broadcast,
                          message: val.target.value,
                        });
                      }}
                      style={emojiStyle}
                      placeholder={
                        "Press ':' for emojis, and template using '@' and shift+enter to send"
                      }
                      autoFocus
                    >
                      <Mention
                        trigger="@"
                        data={[
                          { id: "{{user}}", display: "Full Name" },
                          { id: "{{phone}}", display: "Phone Number" },
                        ]}
                        style={{
                          backgroundColor: "#cee4e5",
                        }}
                        appendSpaceOnAdd
                      />
                      <Mention
                        trigger=":"
                        markup="__id__"
                        regex={neverMatchingRegex}
                        data={queryEmojis}
                      />
                    </MentionsInput>
                  ) : (
                    parseMentions(broadcast.message ?? "", (type, id) => {})
                  )}
                </p>
              </div>
              <div>
                <Label>Status</Label>
                <div className="w-fit">
                  <Badge
                    color={
                      broadcast.status === "DRAFT"
                        ? "warning"
                        : broadcast.status === "FAILED"
                        ? "danger"
                        : "success"
                    }
                  >
                    {broadcast.status}
                  </Badge>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </AdminLayout>
  );
};
export default BroadcastDetail;

const emojiStyle = {
  control: {
    fontSize: 16,
    lineHeight: 1.2,
    minHeight: 160,
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
