import { Modal } from "flowbite-react";
import { useEffect, useState, type FC } from "react";
import { Mention, MentionsInput } from "react-mentions";
import ModalProductList from "./ModalProductList";

interface MessageMentionProps {
  msg: string;
  onChange: (val: any) => void;
  onClickEmoji: () => void;
  onSelectEmoji: (emoji: any) => void;
  children?: React.ReactNode;
}
const neverMatchingRegex = /($a)/;
const MessageMention: FC<MessageMentionProps> = ({
  msg,
  onChange,
  onClickEmoji,
  onSelectEmoji,
  children
}) => {
  const [emojis, setEmojis] = useState<any[]>([]);
  const [modalEmojis, setModalEmojis] = useState(false);
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
    fetch(process.env.REACT_APP_BASE_URL + "/assets/static/emojis.json")
      .then((response) => {
        return response.json();
      })
      .then((jsonData) => {
        setEmojis(jsonData.emojis);
      });
  }, []);
  const groupBy = (emojis: any[], category: string): { [s: string]: any[] } => {
    return emojis.reduce((acc, curr) => {
      const key = curr[category];
      if (!acc[key]) {
        acc[key] = [];
      }
      acc[key].push(curr);
      return acc;
    }, {});
  };
  return (
    <div className="relative w-full">
      <MentionsInput
        value={msg}
        onChange={onChange}
        style={emojiStyle}
        placeholder={"Press ':' for emojis and shift+enter for new line"}
        className="w-full bg-white"
        autoFocus
      >
        <Mention
          trigger="@"
          data={[
            { id: "{{user}}", display: "Full Name" },
            { id: "{{phone}}", display: "Phone Number" },
            { id: "{{agent}}", display: "Agent Name" },
            { id: "{{product}}", display: "Product" },
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
      <div
        className="absolute bottom-1 left-2 cursor-pointer"
        onClick={() => {
          setModalEmojis(true);
          onClickEmoji();
        }}
      >
        😀
      </div>
      {children}
      <Modal
        dismissible
        show={modalEmojis}
        onClose={() => setModalEmojis(false)}
      >
        <Modal.Header>Emojis</Modal.Header>
        <Modal.Body>
          <div>
            {Object.entries(groupBy(emojis, "category")).map(
              ([category, emojis]) => (
                <div
                  className="mb-4 hover:bg-gray-100 rounded-lg p-2"
                  key={category}
                >
                  <h3 className="font-bold">{category}</h3>
                  <div className=" flex flex-wrap gap-1">
                    {emojis.map((e: any, index: number) => (
                      <div
                        key={index}
                        className="cursor-pointer text-lg"
                        onClick={() => {
                         onSelectEmoji(e.emoji);
                          // setModalEmojis(false);
                        }}
                      >
                        {e.emoji}
                      </div>
                    ))}
                  </div>
                </div>
              )
            )}
          </div>
        </Modal.Body>
      </Modal>
      
    </div>
  );
};
export default MessageMention;

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
