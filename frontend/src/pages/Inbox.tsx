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
  getInboxMessageDetail,
  getInboxMessages,
  getInboxMessagesCount,
  getSentMessages,
  getSentMessagesCount,
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
  const [sentUnreadCount, setSentUnreadCount] = useState(0);
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
  const [selectedInbox, setSelectedInbox] = useState<InboxModel>();
  const [selectedMessage, setSelectedMessage] = useState<InboxMessageModel>();
  const [isSentInbox, setIsSentInbox] = useState(false);
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
  useEffect(() => {
    if (mounted) {
      getCount();
    }
  }, [mounted]);
  const getCount = () => {
    getInboxMessagesCount()
      .then((resp: any) => setInboxUnreadCount(resp.data))
      .catch(console.error);
    getSentMessagesCount()
      .then((resp: any) => setSentUnreadCount(resp.data))
      .catch(console.error);
  };
  useEffect(() => {
    if (inbox && mounted) {
      getAllMessages(inbox.id);
      setSelectedInbox(inbox);
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

  const getAllSentMessage = () => {
    getSentMessages({ page, size, search }).then((resp: any) => {
      setMessages(resp.data.items);
      setPagination(getPagination(resp.data));
      setIsSentInbox(true);
    });
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
      setDescription("")
      setSubject("")
      setSelectedMember(null)
      setFiles([])
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  const processReply = async () => {
    if (!selectedMessage) {
      toast.error("Please select message first");
      return;
    }

    try {
      setLoading(true);
      let message = selectedMessage;
      if (selectedMessage.parent) {
        message = selectedMessage.parent;
      }
      var data = {
        recipient_member_id: message!.sender_member_id,
        subject: `[REPLY] ${message?.subject}`,
        message: description,
        attachments: files,
        parent_inbox_message_id: message!.id,
      };
      setSelectedMessage(undefined);
      await sendMessage(data);
      setDescription("")
      setSubject("")
      setSelectedMember(null)
      setFiles([])
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const renderMessages = () => (
    <Table className="table-message" hoverable>
      <Table.Head>
        <Table.HeadCell>
          <Checkbox
            checked={
              messages.every((e) => checkeds.includes(e.id!)) &&
              messages.length > 0
            }
            onClick={(val) => {
              let checked = messages.every((e) => checkeds.includes(e.id!));
              if (!checked) {
                setCheckeds(messages.map((e) => e.id!));
              } else {
                setCheckeds([]);
              }
            }}
          />
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
          <Table.Row style={{ backgroundColor: !e.read ? "#f5f5f5" : "#fff" }}>
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
            <Table.Cell>
              <div
                className="hover:underline cursor-pointer hover:font-semibold"
                onClick={() => {
                  setLoading(true);
                  getInboxMessageDetail(e.parent_inbox_message_id ?? e.id!)
                    .then((resp: any) => {
                      // console.log(resp.data);
                      setSelectedMessage(resp.data);
                      if (selectedInbox) getAllMessages(selectedInbox!.id);
                      if (isSentInbox) {
                        getAllSentMessage();
                      }
                      getCount();
                    })
                    .catch(toast.error)
                    .finally(() => setLoading(false));
                }}
              >
                {e.subject}
              </div>
            </Table.Cell>
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
              {!selectedInbox?.is_trash && !isSentInbox && (
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
                        if (isSentInbox) {
                          getAllSentMessage();
                        } else {
                          getAllMessages(e.inbox_id!);
                        }
                      });
                    }
                  }}
                >
                  Delete
                </a>
              )}
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
          <div className="flex gap-2">
            {checkeds.length > 0 && !selectedInbox?.is_trash && (
              <Button
                color="failure"
                pill
                onClick={async () => {
                  if (
                    window.confirm(
                      `Are you sure you want to delete all messages?`
                    )
                  ) {
                    for (const id of checkeds) {
                      await deleteMessage(id!);
                    }
                    if (isSentInbox) {
                      getAllSentMessage();
                    } else {
                      getAllMessages(inbox?.id!);
                    }
                    setCheckeds([]);
                  }
                }}
              >
                Delete All
              </Button>
            )}

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
        </div>
        <div className="flex flex-row w-full h-full flex-1 gap-2">
          <div className="w-[300px] h-full">
            <ul className="space-y-2">
              <li
                className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold"
                onClick={() => {
                  getAllMessages(inbox!.id);
                  setSelectedInbox(inbox);
                  setIsSentInbox(false);
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
                  getAllSentMessage();
                  setSelectedInbox(undefined);
                }}
              >
                <div className="flex gap-2 items-center">
                  <BsSend /> Sent
                </div>
                {sentUnreadCount > 0 && (
                  <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300">
                    {sentUnreadCount}
                  </span>
                )}
              </li>
              <li
                className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold"
                onClick={() => {
                  getAllMessages(trash!.id);
                  setSelectedInbox(trash);
                  setIsSentInbox(false);
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
      <Drawer
        open={selectedMessage !== undefined}
        onClose={() => setSelectedMessage(undefined)}
        position="right"
        style={{ width: !isFullScreen ? "1000px" : "100%" }}
      >
        <Drawer.Items className="mt-16">
          <h2 className="text-xl font-bold">Message Details</h2>
        </Drawer.Items>
        <Drawer.Items className="mt-4">
          <table className="w-full">
            <tbody>
              <tr>
                <td className="border px-2 py-1 bg-gray-50 font-semibold">
                  Subject
                </td>
                <td className="border px-2 py-1">{selectedMessage?.subject}</td>
              </tr>
              <tr>
                <td className="border px-2 py-1 bg-gray-50 font-semibold">
                  Sender
                </td>
                <td className="border px-2 py-1">
                  <div className="flex  gap-2">
                    <Avatar
                      rounded
                      img={
                        selectedMessage?.sender_member?.user?.profile_picture
                          ?.url
                      }
                      alt="Avatar"
                      size="xs"
                      placeholderInitials={initial(
                        selectedMessage?.sender_member?.user?.full_name
                      )}
                      color="blue"
                    />
                    <h3 className="font-semibold">
                      {selectedMessage?.sender_member?.user?.full_name}
                    </h3>
                  </div>
                </td>
              </tr>
              <tr>
                <td className="border px-2 py-1 bg-gray-50 font-semibold">
                  Date
                </td>
                <td className="border px-2 py-1">
                  <Moment format="DD MMM YYYY, HH:mm">
                    {selectedMessage?.date}
                  </Moment>{" "}
                  (
                  <Moment fromNow className="text-xs">
                    {selectedMessage?.date}
                  </Moment>
                  )
                </td>
              </tr>
            </tbody>
          </table>
        </Drawer.Items>
        {selectedMessage?.parent && (
          <Drawer.Items className="mt-4 ">
            <h3 className="text-lg font-bold">
              Original Message : {selectedMessage?.parent?.subject}
            </h3>

            <div className="pl-8 mt-4">
              <div className="flex flex-col space-y-2 p-4 rounded-lg border">
                <div className="flex justify-between">
                  <div className="flex  gap-2">
                    <Avatar
                      rounded
                      img={
                        selectedMessage?.parent.sender_member?.user
                          ?.profile_picture?.url
                      }
                      alt="Avatar"
                      size="xs"
                      placeholderInitials={initial(
                        selectedMessage?.parent.sender_member?.user?.full_name
                      )}
                      color="blue"
                    />
                    <h3 className="font-semibold">
                      {selectedMessage?.parent?.sender_member?.user?.full_name}
                    </h3>
                  </div>
                  <Moment className="text-sm italic" fromNow>
                    {selectedMessage?.parent?.date}
                  </Moment>
                </div>
                <div
                  className="message-content"
                  dangerouslySetInnerHTML={{
                    __html: selectedMessage?.parent?.message ?? "",
                  }}
                />
                <ul className="">
                  {(selectedMessage?.parent?.attachments ?? []).map(
                    (attachment, index) => (
                      <li key={index}>
                        <a
                          href={attachment?.url}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-blue-500 hover:text-blue-700"
                        >
                          {attachment?.file_name}
                        </a>
                      </li>
                    )
                  )}
                </ul>
              </div>
            </div>
          </Drawer.Items>
        )}
        <Drawer.Items className="mt-4">
          <div
            className="message-content"
            dangerouslySetInnerHTML={{ __html: selectedMessage?.message ?? "" }}
          ></div>
        </Drawer.Items>
        {(selectedMessage?.attachments ?? []).length > 0 && (
          <Drawer.Items className="mt-4">
            <h3 className="text-lg font-bold">Attachments</h3>
            <ul className="list-disc list-inside">
              {(selectedMessage?.attachments ?? []).map((attachment, index) => (
                <li key={index}>
                  <a
                    href={attachment?.url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-blue-500 hover:text-blue-700"
                  >
                    {attachment?.file_name}
                  </a>
                </li>
              ))}
            </ul>
          </Drawer.Items>
        )}
        {(selectedMessage?.replies ?? []).length > 0 && (
          <Drawer.Items className="mt-4 ">
            <h3 className="text-lg font-bold">Replies</h3>

            <div className="pl-8 mt-4 space-y-4">
              {(selectedMessage?.replies ?? []).map((reply, index) => (
                <div
                  key={index}
                  className="flex flex-col space-y-2 p-4 rounded-lg border"
                >
                  <div className="flex justify-between">
                    <div className="flex  gap-2">
                      <Avatar
                        rounded
                        img={reply.sender_member?.user?.profile_picture?.url}
                        alt="Avatar"
                        size="xs"
                        placeholderInitials={initial(
                          reply.sender_member?.user?.full_name
                        )}
                        color="blue"
                      />
                      <h3 className="font-semibold">
                        {reply?.sender_member?.user?.full_name}
                      </h3>
                    </div>
                    <Moment className="text-sm italic" fromNow>
                      {reply?.date}
                    </Moment>
                  </div>
                  <div
                    className="message-content"
                    dangerouslySetInnerHTML={{ __html: reply?.message ?? "" }}
                  />
                  <ul className="">
                    {(reply?.attachments ?? []).map((attachment, index) => (
                      <li key={index}>
                        <a
                          href={attachment?.url}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-blue-500 hover:text-blue-700"
                        >
                          {attachment?.file_name}
                        </a>
                      </li>
                    ))}
                  </ul>
                </div>
              ))}
            </div>
          </Drawer.Items>
        )}
        <FooterDivider />
        <Drawer.Items style={{ minHeight: 400 }}>
          <h3 className="text-lg font-bold">Reply</h3>
          <Editor
            apiKey={process.env.REACT_APP_TINY_MCE_KEY}
            init={{
              height: 400,
              plugins:
                "anchor autolink charmap codesample emoticons image link lists media searchreplace table visualblocks wordcount ",
              toolbar:
                "undo redo | blocks fontfamily fontsize | bold italic underline strikethrough | forecolor backcolor | link image media table | align lineheight | numlist bullist indent outdent | emoticons charmap | removeformat",
            }}
            initialValue={description ?? ""}
            onChange={handleEditorChange}
          />
        </Drawer.Items>
        <Drawer.Items className="mt-12">
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
            <Button color="blue" onClick={processReply}>
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
