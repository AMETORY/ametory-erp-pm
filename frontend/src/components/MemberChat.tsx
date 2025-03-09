import { useEffect, useState, type FC } from "react";
import { ProjectModel } from "../models/project";
import { Avatar } from "flowbite-react";
import { initial, isEmailFormatValid } from "../utils/helper";
import { HiMiniMagnifyingGlass } from "react-icons/hi2";
import { getMembers, inviteMember } from "../services/api/commonApi";
import { MemberModel } from "../models/member";
import { getProjectAddMember, getProjectMembers } from "../services/api/projectApi";
import toast from "react-hot-toast";
import { ChatChannelModel } from "../models/chat";

interface MemberChatModalProps {
  channel: ChatChannelModel;
  onInvite: (v: string) => void
}

const MemberChatModal: FC<MemberChatModalProps> = ({ channel, onInvite }) => {
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(10);
  const [search, setSearch] = useState("");
  const [searchExisting, setSearchExisting] = useState("");
  const [members, setMembers] = useState<MemberModel[]>([]);
  const [existingMembers, setExistingMembers] = useState<MemberModel[]>([]);
  useEffect(() => {}, []);

  useEffect(() => {
    getMembers({ page, size, search })
      .then((res: any) => {
        setMembers(res.data.items);
      })
      .catch(toast.error);

      getAllMembers()
  }, [page, size, search]);

  useEffect(() => {
    setExistingMembers(channel.participant_members ?? [])
  }, []);

  const renderUser = (member: MemberModel) => (
    <div className="flex flex-row gap-2">
      <Avatar
        size="xs"
        img={member?.user?.profile_picture?.url}
        rounded
        stacked
        placeholderInitials={initial(member?.user?.full_name)}
      />
      <div className="flex flex-col hover:font-semibold">
        <span className="">
          {member.user?.full_name}
        </span>
        <small className="">{member.user?.email}</small>
      </div>
    </div>
  );

  const getAllMembers = () => {
 

  }
  return (
    <div className="flex gap-4">
      <div className="w-1/2 ">
        <h2 className="text-lg font-bold text-gray-500">Search Member</h2>
        <div className="flex items-center mt-2 relative">
          <input
            type="search"
            className="border-gray-400 outline-2 rounded-lg px-2 py-1 w-full pr-8 "
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            placeholder="Search member or add by email"
          />
          <HiMiniMagnifyingGlass className="absolute right-2" />
        </div>
        {members.length == 0 &&
          search.length > 3 &&
          isEmailFormatValid(search) && (
            <small className="px-2 w-full cursor-pointer hover:bg-gray-100" onClick={() => onInvite(search)}>Invite "{search}" as a member</small>
          )}
        <ul className="mb-2 mt-4">
          {(members ?? []).filter(e => !(existingMembers).map(m => m.id).includes(e.id)).map((member) => (
            <li
              key={member.id}
              className="py-2 border-b last:border-b-0 flex cursor-pointer"
              onClick={() => {
                
              }}
            >
              {renderUser(member)}
            </li>
          ))}
        </ul>
      </div>
      <div className="w-1/2">
        <h2 className="text-lg font-bold text-gray-500">Existing Members</h2>
        <div className="flex items-center mt-2 relative">
          <input
            type="search"
            className="border-gray-400 outline-2 rounded-lg px-2 py-1 w-full pr-8 "
            value={searchExisting}
            onChange={(e) => setSearchExisting(e.target.value)}
            placeholder="Search member"
          />
          <HiMiniMagnifyingGlass className="absolute right-2" />
        </div>
        <ul className="mb-2 mt-4">
          {(existingMembers)
            .filter((member) =>
              member.user.full_name
                ?.toLowerCase()
                .includes(searchExisting.toLowerCase())
            )
            .map((member) => (
              <li
                key={member.id}
                className="py-2 border-b last:border-b-0 flex cursor-pointer"
              >
                {renderUser(member)}
              </li>
            ))}
        </ul>
      </div>
    </div>
  );
};
export default MemberChatModal;
