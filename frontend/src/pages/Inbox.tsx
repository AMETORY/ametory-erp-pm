import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Avatar,
  Button,
  Checkbox,
  Drawer,
  FileInput,
  FooterDivider,
  Label,
  Table,
  TextInput,
} from "flowbite-react";
import {
  deleteMessage,
  getInboxes,
  getInboxMessages,
  sendMessage,
} from "../services/api/inboxApi";
import { InboxMessageModel, InboxModel } from "../models/inbox";
import { PaginationResponse } from "../objects/pagination";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";
import { HiOutlineInboxArrowDown, HiOutlineTrash } from "react-icons/hi2";
import { HiOutlineDocument, HiPaperAirplane, HiTrash } from "react-icons/hi";
import { getPagination, initial } from "../utils/helper";
import { MemberModel } from "../models/member";
import { getMembers, uploadFile } from "../services/api/commonApi";
import Select, { InputActionMeta } from "react-select";
import { Editor } from "@tinymce/tinymce-react";
import { BsSend } from "react-icons/bs";
import { FileModel } from "../models/file";
import { WebsocketContext } from "../contexts/WebsocketContext";
import Moment from "react-moment";

interface InboxPageProps {}

const InboxPage: FC<InboxPageProps> = ({}) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { loading, setLoading } = useContext(LoadingContext);
  const [mounted, setMounted] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(10);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [inbox, setInbox] = useState<InboxModel>();
  const [trash, setTrash] = useState<InboxModel>();
  const [inboxUnreadCount, setInboxUnreadCount] = useState(0);
  const [messages, setMessages] = useState<InboxMessageModel[]>([]);
  const [openCompose, setOpenCompose] = useState(false);
  const [isFullScreen, setIsFullScreen] = useState(false);
  const [members, setMembers] = useState<MemberModel[]>([]);
  const [selectedMember, setSelectedMember] = useState<{
    label: string;
    value: string;
  } | null>(null);
  const [subject, setSubject] = useState("");
  const [description, setDescription] = useState("");
  const [files, setFiles] = useState<FileModel[]>([]);
  const [checkeds, setCheckeds] = useState<string[]>([]);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getInboxes().then((resp: any) => {
        setInbox(resp.data.find((e: InboxModel) => e.is_default));
        setTrash(resp.data.find((e: InboxModel) => e.is_trash));
      });
      getAllMember("");
    }
  }, [mounted]);
  useEffect(() => {}, []);
  useEffect(() => {
    if (inbox && mounted) {
      getAllMessages(inbox.id);
    }
  }, [inbox, mounted, page, size, search]);
  const getAllMessages = async (inboxID: string) => {
    try {
      setLoading(true);
      let resp: any = await getInboxMessages(inboxID, { page, size, search });
      setMessages(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (inbox) {
      if (wsMsg?.inbox_id == inbox!.id) {
        toast.success("You have new message");
        getAllMessages(wsMsg?.inbox_id!);
      }
    }
  }, [inbox, wsMsg]);

  const getAllMember = (s: string) => {
    getMembers({ page: 1, size: 10, search: s })
      .then((res: any) => {
        setMembers(res.data.items);
      })
      .catch(toast.error);
  };

  const handleEditorChange = (e: any) => {
    setDescription(e.target.getContent());
  };

  const processSend = async () => {
    if (subject.length == 0) {
      toast.error("Please enter subject");
      return;
    }
    if (!selectedMember) {
      toast.error("Please select recipient");
      return;
    }
    try {
      setLoading(true);
      var data = {
        recipient_member_id: selectedMember!.value,
        subject: subject,
        message: description,
        attachments: files,
      };
      setOpenCompose(false);
      await sendMessage(data);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const renderMessages = () => (
    <Table className="table-message">
      <Table.Head>
        <Table.HeadCell>
          <Checkbox />
        </Table.HeadCell>
        <Table.HeadCell>Sender</Table.HeadCell>
        <Table.HeadCell>Subject</Table.HeadCell>
        <Table.HeadCell>Date</Table.HeadCell>
        <Table.HeadCell></Table.HeadCell>
      </Table.Head>
      <Table.Body>
        {messages.length === 0 && (
          <Table.Row>
            <Table.Cell colSpan={5} className="text-center">
              No messages found.
            </Table.Cell>
          </Table.Row>
        )}
      </Table.Body>
      {messages.map((e) => (
        <Table.Body key={e.id}>
          <Table.Row>
            <Table.Cell width={20}>
              <Checkbox
                checked={checkeds.includes(e.id!)}
                onChange={(val) => {
                  if (val.target.checked) {
                    setCheckeds((val) => [...val, e.id!]);
                  } else {
                    setCheckeds(checkeds.filter((id) => id !== e.id));
                  }
                }}
              />
            </Table.Cell>
            <Table.Cell width={200}>
              {e.sender_member?.user?.full_name}
            </Table.Cell>
            <Table.Cell>{e.subject}</Table.Cell>
            <Table.Cell width={160}>
              <Moment format="DD MMM YYYY">{e.date}</Moment>
            </Table.Cell>
            <Table.Cell width={160}>
              {/* <a
                                href="#"
                                className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                              >
                                Edit
                              </a> */}
              <a
                href="#"
                className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                onClick={(el) => {
                  el.preventDefault();
                  if (
                    window.confirm(
                      `Are you sure you want to delete project ${e.subject}?`
                    )
                  ) {
                    deleteMessage(e?.id!).then(() => {
                      getAllMessages(e.inbox_id!);
                    });
                  }
                }}
              >
                Delete
              </a>
            </Table.Cell>
          </Table.Row>
        </Table.Body>
      ))}
    </Table>
  );
  return (
    <AdminLayout>
      <div className="p-4 flex flex-col h-full">
        <div className="flex justify-between items-center mb-2 border-b pb-4">
          <h1 className="text-3xl font-bold ">Inbox</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setOpenCompose(true);
            }}
          >
            + Compose
          </Button>
        </div>
        <div className="flex flex-row w-full h-full flex-1 gap-2">
          <div className="w-[300px] h-full">
            <ul className="space-y-2">
              <li
                className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold"
                onClick={() => {
                  getAllMessages(inbox!.id);
                }}
              >
                <div className="flex gap-2 items-center">
                  <HiOutlineInboxArrowDown /> Inbox
                </div>
                {inboxUnreadCount > 0 && (
                  <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300">
                    {inboxUnreadCount}
                  </span>
                )}
              </li>
              <li
                className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold"
                onClick={() => {
                  getAllMessages(trash!.id);
                }}
              >
                <div className="flex gap-2 items-center">
                  <HiOutlineTrash /> Trash
                </div>
              </li>
            </ul>
          </div>
          <div className="w-full border-l px-4">{renderMessages()}</div>
        </div>
      </div>
      <Drawer
        position="right"
        style={{ width: !isFullScreen ? "1000px" : "100%" }}
        open={openCompose}
        onClose={() => setOpenCompose(false)}
      >
        <Drawer.Header>
          <h1 className="text-2xl font-bold">Compose</h1>
        </Drawer.Header>

        <Drawer.Items>
          <div className="flex flex-col mt-8 space-y-4">
            <div className="flex items-center">
              <div className="w-[160px]">To</div>
              <div className="flex-1">
                <Select
                  className="w-full"
                  value={selectedMember}
                  onChange={(val) => {
                    setSelectedMember(val);
                  }}
                  options={(members ?? []).map((member) => ({
                    label: member.user?.full_name ?? "",
                    value: member.id ?? "",
                    avatar: (
                      <Avatar
                        rounded
                        img={member?.user?.profile_picture?.url}
                        alt="Avatar"
                        size="xs"
                        placeholderInitials={initial(member?.user?.full_name)}
                        color="blue"
                      />
                    ),
                  }))}
                  formatOptionLabel={(option: any) => (
                    <div className="flex flex-row gap-2  items-center">
                      {option.avatar}
                      <span>{option.label}</span>
                    </div>
                  )}
                  inputValue={""}
                  onInputChange={(
                    newValue: string,
                    actionMeta: InputActionMeta
                  ) => {
                    // console.log(newValue, actionMeta);
                  }}
                />
              </div>
            </div>
            <div className="flex items-center">
              <div className="w-[160px]">Subject</div>
              <div className="flex-1">
                <TextInput
                  value={subject}
                  onChange={(val) => setSubject(val.target?.value)}
                />
              </div>
            </div>
          </div>
        </Drawer.Items>
        <FooterDivider />
        <Drawer.Items style={{ height: 600 }}>
          <Editor
            apiKey={process.env.REACT_APP_TINY_MCE_KEY}
            init={{
              height: 600,
              plugins:
                "anchor autolink charmap codesample emoticons image link lists media searchreplace table visualblocks wordcount ",
              toolbar:
                "undo redo | blocks fontfamily fontsize | bold italic underline strikethrough | forecolor backcolor | link image media table | align lineheight | numlist bullist indent outdent | emoticons charmap | removeformat",
            }}
            initialValue={description ?? ""}
            onChange={handleEditorChange}
          />
        </Drawer.Items>
        <Drawer.Items className="mt-4">
          {files.length == 0 ? (
            <div className="flex w-full items-center justify-center">
              <Label
                htmlFor="dropzone-file"
                className="flex h-[160px] w-full  cursor-pointer flex-col items-center justify-center rounded-lg border-2 border-dashed border-gray-300 bg-gray-50 hover:bg-gray-100 dark:border-gray-600 dark:bg-gray-700 dark:hover:border-gray-500 dark:hover:bg-gray-600"
              >
                <div className="flex flex-col items-center justify-center py-2">
                  <p className="mb-2 text-sm text-gray-500 dark:text-gray-400">
                    <span className="font-semibold">Click to upload</span> or
                    drag and drop
                  </p>
                  <p className="text-xs text-gray-500 dark:text-gray-400">
                    Document or Image
                  </p>
                </div>
                <FileInput
                  onChange={(val) => {
                    const files = val?.target.files;
                    if (files) {
                      for (let index = 0; index < files.length; index++) {
                        const file = files[index];
                        try {
                          uploadFile(file, {}, (val) => console.log).then(
                            (v: any) => {
                              setFiles((files) => [...files, v.data]);
                            }
                          );
                        } catch (error) {
                          console.log(error);
                        }
                      }
                    }
                  }}
                  multiple
                  id="dropzone-file"
                  className="hidden"
                  accept="image/*,.doc,.docx,.pdf,.xlsx,.xls,.csv"
                />
              </Label>
            </div>
          ) : (
            <ul>
              {files.map((e) => (
                <li
                  key={e.id}
                  className="cursor-pointer flex gap-2 items-center"
                  onClick={() => {
                    window.open(e.url);
                  }}
                >
                  <HiOutlineDocument /> {e.file_name}
                </li>
              ))}
            </ul>
          )}
        </Drawer.Items>
        <Drawer.Items className="mt-8">
          <div className="flex justify-end">
            <Button color="blue" onClick={processSend}>
              <BsSend className="mr-2" />
              Send
            </Button>
          </div>
        </Drawer.Items>
      </Drawer>
    </AdminLayout>
  );
};
export default InboxPage;
