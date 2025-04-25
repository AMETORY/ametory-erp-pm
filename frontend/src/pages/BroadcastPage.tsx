import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Badge,
  Button,
  Label,
  Modal,
  Textarea,
  TextInput,
} from "flowbite-react";
import { Mention, MentionsInput } from "react-mentions";
import { PaginationResponse } from "../objects/pagination";
import { createBroadcast, getBroadcasts } from "../services/api/broadcastApi";
import { SearchContext } from "../contexts/SearchContext";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";

interface BroadcastPageProps {}
const neverMatchingRegex = /($a)/;
const BroadcastPage: FC<BroadcastPageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [mounted, setMounted] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [emojis, setEmojis] = useState([]);
  const [message, setMessage] = useState("");
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const { search, setSearch } = useContext(SearchContext);
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [description, setDescription] = useState("");
  const nav = useNavigate();
  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getBroadcasts({ page, size, search }).then((res) => {});
    }
  }, [mounted, page, size, search]);
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

  const send = async () => {
    try {
      if (description.trim().length === 0) {
        toast.error("description is required");
        return;
      }
      if (message.trim().length === 0) {
        toast.error("Message is required");
        return;
      }
      setLoading(true);
      let resp: any = await createBroadcast({ message, description });
      nav(`/broadcast/${resp.data.id}`);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Broadcast</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setShowModal(true);
            }}
          >
            + Create new broadcast
          </Button>
        </div>
      </div>
      {showModal && (
        <Modal show={showModal} onClose={() => setShowModal(false)}>
          <Modal.Header>Create new broadcast</Modal.Header>
          <Modal.Body>
            <form>
              <div className="mb-2 block">
                <Label htmlFor="description" value="Description" />
                <Textarea
                  id="description"
                  placeholder="Input description"
                  className="mt-1 input-white"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                />
              </div>
              <div className="mb-2 block">
                <Label htmlFor="message" value="Message" />

                <MentionsInput
                  value={message}
                  onChange={(val) => {
                    setMessage(val.target.value);
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
              </div>
            </form>
          </Modal.Body>
          <Modal.Footer>
            <div className="flex w-full justify-end">
              <Button type="submit" color="success" onClick={send}>
                Save
              </Button>
            </div>
          </Modal.Footer>
        </Modal>
      )}
    </AdminLayout>
  );
};
export default BroadcastPage;

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
