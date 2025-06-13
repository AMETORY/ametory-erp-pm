import { act, useContext, useEffect, useState, type FC } from "react";
import { LoadingContext } from "../contexts/LoadingContext";
import { MemberInvitationModel, MemberModel } from "../models/member";
import { PaginationResponse } from "../objects/pagination";
import {
  deleteInvitation,
  getInvitedMembers,
  getMembers,
  getRoles,
  inviteMember,
  updateMember,
} from "../services/api/commonApi";
import { getPagination, initial } from "../utils/helper";
import toast from "react-hot-toast";
import AdminLayout from "../components/layouts/admin";
import {
  Avatar,
  Badge,
  Button,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  Pagination,
  Table,
  Tabs,
  TextInput,
} from "flowbite-react";
import { useNavigate } from "react-router-dom";
import { RoleModel } from "../models/role";
import { MemberContext, ProfileContext } from "../contexts/ProfileContext";

interface MemberPageProps {}

const MemberPage: FC<MemberPageProps> = ({}) => {
  const { profile } = useContext(ProfileContext);
  const { member: activeMember } = useContext(MemberContext);
  const [inviteModal, setInviteModal] = useState(false);
  const { loading, setLoading } = useContext(LoadingContext);
  const [members, setMembers] = useState<MemberModel[]>([]);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [member, setMember] = useState<MemberModel | null>(null);
  const navigate = useNavigate();
  const [inviteEmail, setInviteEmail] = useState("");
  const [inviteFullName, setInviteFullName] = useState("");
  const [inviteRoleId, setInviteRoleId] = useState("");
  const [roles, setRoles] = useState<RoleModel[]>([]);

  const [pageInvited, setPageInvited] = useState(1);
  const [sizeInvited, setSizeInvited] = useState(20);
  const [searchInvited, setSearchInvited] = useState("");
  const [paginationInvited, setPaginationInvited] =
    useState<PaginationResponse>();
  const [inviteds, setInviteds] = useState<MemberInvitationModel[]>([]);

  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllMembers();
    }
  }, [mounted, page, size, search]);
  useEffect(() => {
    if (mounted) {
      getAllInvited();
    }
  }, [mounted, pageInvited, sizeInvited, searchInvited]);
  useEffect(() => {
    getRoles({ page: 1, size: 10, search: "" }).then((res: any) => {
      setRoles(res.data.items);
    });
  }, []);

  const getAllMembers = async () => {
    try {
      setLoading(true);
      const resp: any = await getMembers({ page, size, search });
      setMembers(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  const getAllInvited = async () => {
    try {
      setLoading(true);
      const resp: any = await getInvitedMembers({
        page: pageInvited,
        size: sizeInvited,
        search: searchInvited,
      });
      setInviteds(resp.data.items);
      setPaginationInvited(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const renderInvited = () => (
    <div>
      <Table>
        <Table.Head>
          <Table.HeadCell>Name</Table.HeadCell>
          <Table.HeadCell>Email</Table.HeadCell>
          <Table.HeadCell>Role</Table.HeadCell>
          <Table.HeadCell>Invited By</Table.HeadCell>
          <Table.HeadCell></Table.HeadCell>
        </Table.Head>
        <Table.Body className="divide-y">
          {inviteds.length === 0 && (
            <Table.Row>
              <Table.Cell colSpan={5} className="text-center">
                No member invitation found.
              </Table.Cell>
            </Table.Row>
          )}
          {inviteds.map((invited) => (
            <Table.Row
              key={invited.id}
              className="bg-white dark:border-gray-700 dark:bg-gray-800"
            >
              <Table.Cell
                className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold"
                onClick={() => {}}
              >
                {invited.full_name}
              </Table.Cell>
              <Table.Cell>{invited?.email}</Table.Cell>
              <Table.Cell>
                <div className="flex">
                  {invited?.role?.is_super_admin && (
                    <Badge color="green">{invited?.role?.name}</Badge>
                  )}
                  {invited?.role?.is_admin && (
                    <Badge>{invited?.role?.name}</Badge>
                  )}
                  {!invited?.role?.is_admin &&
                    !invited?.role?.is_super_admin && (
                      <Badge color="pink">{invited?.role?.name}</Badge>
                    )}
                </div>
              </Table.Cell>
              <Table.Cell>{invited?.inviter?.full_name}</Table.Cell>
              <Table.Cell>
                {/* <a
                    href="#"
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                  >
                    Edit
                  </a> */}
                <a
                  href="#"
                  className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                  onClick={(e) => {
                    e.preventDefault();
                    if (
                      window.confirm(
                        `Are you sure you want to delete invitation ${invited.full_name}?`
                      )
                    ) {
                      deleteInvitation(invited?.id!).then(() => {
                        getAllInvited();
                      });
                    }
                  }}
                >
                  Delete
                </a>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </div>
  );
  const renderMembers = () => (
    <div>
      <Table>
        <Table.Head>
          <Table.HeadCell>Name</Table.HeadCell>
          <Table.HeadCell>Email</Table.HeadCell>
          <Table.HeadCell>Role</Table.HeadCell>
          <Table.HeadCell>Avatar</Table.HeadCell>
          <Table.HeadCell></Table.HeadCell>
        </Table.Head>

        <Table.Body className="divide-y">
          {members.length === 0 && (
            <Table.Row>
              <Table.Cell colSpan={5} className="text-center">
                No members found.
              </Table.Cell>
            </Table.Row>
          )}
          {members.map((member) => (
            <Table.Row
              key={member.id}
              className="bg-white dark:border-gray-700 dark:bg-gray-800"
            >
              <Table.Cell
                className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold"
                onClick={() => navigate(`/member/${member.id}`)}
              >
                {member.user?.full_name}
              </Table.Cell>
              <Table.Cell>{member?.user?.email}</Table.Cell>
              <Table.Cell>
                <div className="flex">
                  {member?.role?.is_super_admin && (
                    <Badge color="green">{member?.role?.name}</Badge>
                  )}
                  {member?.role?.is_admin && (
                    <Badge>{member?.role?.name}</Badge>
                  )}
                  {!member?.role?.is_admin && !member?.role?.is_super_admin && (
                    <Badge color="pink">{member?.role?.name}</Badge>
                  )}
                </div>
              </Table.Cell>
              <Table.Cell>
                <Avatar
                  key={member.id}
                  size="xs"
                  img={member?.user?.profile_picture?.url}
                  rounded
                  stacked
                  placeholderInitials={initial(member?.user?.full_name)}
                />
              </Table.Cell>
              <Table.Cell>
                {activeMember?.role?.is_super_admin && (
                  <a
                    href="#"
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                    onClick={(e) => {
                      e.preventDefault();
                      setMember(member);
                    }}
                  >
                    Edit
                  </a>
                )}

                <a
                  href="#"
                  className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                  onClick={(e) => {
                    e.preventDefault();
                    if (
                      window.confirm(
                        `Are you sure you want to delete member ${member.user?.full_name}?`
                      )
                    ) {
                      // deleteMember(member?.id!).then(() => {
                      //   getAllMembers();
                      // });
                    }
                  }}
                >
                  Delete
                </a>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
      <Pagination
        className="mt-4"
        currentPage={page}
        totalPages={pagination?.total_pages ?? 0}
        onPageChange={(val) => {
          setPage(val);
        }}
        showIcons
      />
    </div>
  );
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Member</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setInviteModal(true);
            }}
          >
            + Invite Member
          </Button>
        </div>
        <Tabs aria-label="Pills" variant="pills">
          <Tabs.Item active title="Members">
            {renderMembers()}
          </Tabs.Item>
          <Tabs.Item title="Invited">{renderInvited()}</Tabs.Item>
        </Tabs>
      </div>
      <Modal
        size="4xl"
        show={inviteModal}
        onClose={() => setInviteModal(false)}
      >
        <Modal.Header>Invite Member</Modal.Header>
        <Modal.Body>
          <form className="flex flex-col gap-4">
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700"
              >
                Email
              </label>
              <TextInput
                id="email"
                type="email"
                placeholder="Email"
                value={inviteEmail}
                onChange={(e) => setInviteEmail(e.target.value)}
              />
            </div>
            <div>
              <label
                htmlFor="full_name"
                className="block text-sm font-medium text-gray-700"
              >
                Full Name
              </label>
              <TextInput
                id="full_name"
                type="text"
                placeholder="Full Name"
                value={inviteFullName}
                onChange={(e) => {
                  setInviteFullName(e.target.value);
                }}
              />
            </div>
            <div>
              <label
                htmlFor="role"
                className="block text-sm font-medium text-gray-700"
              >
                Role
              </label>
              <select
                value={inviteRoleId}
                id="role"
                className="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                onChange={(e) => setInviteRoleId(e.target.value)}
              >
                <option value="">Select Role</option>
                {roles.map((role) => (
                  <option key={role.id} value={role.id}>
                    {role.name}
                  </option>
                ))}
              </select>
            </div>
          </form>
        </Modal.Body>
        <Modal.Footer className="flex justify-end">
          <Button
            type="submit"
            color="blue"
            onClick={async () => {
              try {
                if (!inviteFullName || !inviteRoleId) {
                  throw new Error("Please fill in all fields");
                }
                setLoading(true);
                // saveInvite();
                let invitationData = {
                  full_name: inviteFullName,
                  role_id: inviteRoleId,
                  email: inviteEmail,
                };
                await inviteMember(invitationData);
                setInviteModal(false);
                toast.success(" Invite sent successfully");
              } catch (error: any) {
                toast.error(error.message);
              } finally {
                setLoading(false);
              }
            }}
          >
            Save
          </Button>
          <Button color="gray" onClick={() => setInviteModal(false)}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
      <Modal show={member != null} onClose={() => setMember(null)}>
        <ModalHeader>Update Member</ModalHeader>
        <ModalBody>
          <form className="flex flex-col gap-4">
            <div>
              <label
                htmlFor="full_name"
                className="block text-sm font-medium text-gray-700"
              >
                Full Name
              </label>
              <div>{member?.user?.full_name}</div>
            </div>
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700"
              >
                Email
              </label>
              <div>{member?.user?.email}</div>
            </div>
            <div>
              <label
                htmlFor="role"
                className="block text-sm font-medium text-gray-700"
              >
                Role
              </label>
              <select
                value={member?.role?.id}
                id="role"
                className="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                onChange={(e) => {
                  setMember({
                    ...member!,
                    role: {
                      ...member?.role!,
                      id: e.target.value,
                    },
                    role_id: e.target.value,
                  });
                }}
              >
                <option value="">Select Role</option>
                {roles.map((role) => (
                  <option key={role.id} value={role.id}>
                    {role.name}
                  </option>
                ))}
              </select>
            </div>
          </form>
        </ModalBody>
        <ModalFooter className="flex justify-end">
          <Button
            type="submit"
            color="blue"
            onClick={async () => {
              try {
                setLoading(true);
                await updateMember(member?.id!, member);
                setMember(null);
                toast.success("Member updated successfully");
                getAllMembers();
              } catch (error) {
                toast.error(`${error}`);
              } finally {
                setLoading(false);
              }
            }}
          >
            Save
          </Button>
          <Button color="gray" onClick={() => setMember(null)}>
            Close
          </Button>
        </ModalFooter>
      </Modal>
    </AdminLayout>
  );
};
export default MemberPage;
