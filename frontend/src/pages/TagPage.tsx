import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Button,
  Label,
  Modal,
  Pagination,
  Table,
  Textarea,
  TextInput,
} from "flowbite-react";
import { LoadingContext } from "../contexts/LoadingContext";
import { PaginationResponse } from "../objects/pagination";
import {
  createTag,
  deleteTag,
  getTags,
  updateTag,
} from "../services/api/tagApi";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { TagModel } from "../models/tag";
import { getContrastColor, getPagination } from "../utils/helper";

interface TagPageProps {}

const TagPage: FC<TagPageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [showModal, setShowModal] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [name, setName] = useState("");
  const [color, setColor] = useState("");
  const nav = useNavigate();
  const [selectedTag, setSelectedTag] = useState<TagModel>();
  const [mounted, setMounted] = useState(false);
  const [tags, setTags] = useState<TagModel[]>([]);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllTags();
    }
  }, [mounted, page, size, search]);

  const getAllTags = async () => {
    try {
      setLoading(true);
      let resp: any = await getTags({ page, size, search });
      setTags(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  const save = async () => {
    setLoading(true);
    try {
      if (selectedTag) {
        await updateTag(selectedTag!.id, {
          name,
          color,
        });
      } else {
        await createTag({
          name,
          color,
        });
      }
      setShowModal(false);
      setName("");
      setColor("");
      setSelectedTag(undefined);
      toast.success("Save successfully");
      getAllTags();
    } catch (error) {
      console.log(error);
      toast.error("Save failed");
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Tag</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setShowModal(true);
            }}
          >
            + Create new tag
          </Button>
        </div>
        <Table hoverable={true}>
          <Table.Head>
            <Table.HeadCell>Name</Table.HeadCell>
            <Table.HeadCell>Color</Table.HeadCell>
            <Table.HeadCell></Table.HeadCell>
          </Table.Head>
          <Table.Body className="bg-white">
            {tags.map((tag) => (
              <Table.Row
                key={tag.id}
                className="bg-white dark:border-gray-700 dark:bg-gray-800"
              >
                <Table.Cell>
                  <span className="font-medium">{tag.name}</span>
                </Table.Cell>
                <Table.Cell>
                    <div className="px-2 py-1 rounded-lg w-fit" style={{ backgroundColor: tag.color, color: getContrastColor(tag.color) }}>
                    {tag.color}
                    </div>
                </Table.Cell>
                <Table.Cell>
                  <a
                    href="#"
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                    onClick={() => {
                      setSelectedTag(tag);
                      setShowModal(true);
                      setName(tag.name);
                      setColor(tag.color ?? "#ff0000");
                    }}
                  >
                    Edit
                  </a>
                  <a
                    href="#"
                    className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                    onClick={(e) => {
                      e.preventDefault();
                      if (
                        window.confirm(
                          `Are you sure you want to delete  ${tag.name}?`
                        )
                      ) {
                        deleteTag(tag.id).then(() => {
                          getAllTags();
                        })
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
      <Modal show={showModal} onClose={() => setShowModal(false)}>
        <Modal.Header>Create new tag</Modal.Header>
        <Modal.Body>
          <div className="flex flex-col space-y-4">
            <div>
              <Label htmlFor="name" value="Name" className="mb-1" />
              <TextInput
                id="name"
                type="text"
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Name"
                className="mb-4"
              />
            </div>

            <div>
              <Label htmlFor="color" value="Color" className="mb-1" />
              <TextInput
                id="color"
                type="color"
                value={color}
                onChange={(e) => setColor(e.target.value)}
                placeholder="Color"
                className="mb-4"
              />
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="flex flex-row justify-end w-full">
            <Button onClick={save}>Save</Button>
          </div>
        </Modal.Footer>
      </Modal>
    </AdminLayout>
  );
};
export default TagPage;
