import { useContext, useEffect, useRef, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { PaginationResponse } from "../objects/pagination";
import {
  createContact,
  deleteContact,
  getContacts,
  importContact,
  updateContact,
} from "../services/api/contactApi";
import { ContactModel } from "../models/contact";
import { LoadingContext } from "../contexts/LoadingContext";
import { getContrastColor, getPagination, randomColor } from "../utils/helper";
import toast from "react-hot-toast";
import {
  Button,
  Drawer,
  DrawerItems,
  Label,
  Modal,
  Pagination,
  Table,
  Textarea,
  TextInput,
} from "flowbite-react";
import { SearchContext } from "../contexts/SearchContext";
import { TagModel } from "../models/tag";
import { createTag, getTags } from "../services/api/tagApi";
import Select from "react-select";
import CreatableSelect from "react-select/creatable";
import { ProductModel } from "../models/product";
import { getProducts } from "../services/api/productApi";
import { PiFileXls } from "react-icons/pi";
import { uploadFile } from "../services/api/commonApi";
import { BsFilter, BsTag } from "react-icons/bs";
import { LuFilter } from "react-icons/lu";

interface ContactPageProps {}

const ContactPage: FC<ContactPageProps> = ({}) => {
  const { search, setSearch } = useContext(SearchContext);
  const { loading, setLoading } = useContext(LoadingContext);
  const [mounted, setMounted] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [contacts, setContacts] = useState<ContactModel[]>([]);
  const [showModal, setShowModal] = useState(false);
  const [selectedContact, setSelectedContact] = useState<ContactModel>();
  const [products, setProducts] = useState<ProductModel[]>([]);
  const fileRef = useRef<HTMLInputElement>(null);
  const [drawerFilter, setDrawerFilter] = useState(false);
  const [tags, setTags] = useState<TagModel[]>([]);
  const [selectedTags, setSelectedTags] = useState<TagModel[]>([]);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllContacts();
    }
  }, [mounted, page, size, search, selectedTags]);

  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files == null) return;
    const file = e.target.files[0];
    setLoading(true);
    try {
      const resp: any = await uploadFile(
        file,
        {
          skip_save: true,
        },
        (progress) => {
          console.log(progress);
        }
      );
      await importContact({
        file_url: resp.data.url,
      });
    } catch (e) {
      toast.error(`${e}`);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (mounted) {
      getAllTags();
      getAllProducts();
    }
  }, [mounted]);
  const getAllContacts = async () => {
    try {
      setLoading(true);
      let resp: any = await getContacts({
        page,
        size,
        search,
        order: "name",
        tag_ids: selectedTags.map((t) => t.id).join(","),
      });
      setContacts(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const getAllTags = async () => {
    try {
      setLoading(true);
      let resp: any = await getTags({ page: 1, size: 100 });
      setTags(resp.data.items);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const getAllProducts = async () => {
    try {
      setLoading(true);
      let resp: any = await getProducts({ page: 1, size: 100 });
      setProducts(resp.data.items);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const handleCreateContact = async () => {
    try {
      if (!selectedContact) return;
      setLoading(true);
      if (selectedContact?.id) {
        await updateContact(selectedContact!.id, selectedContact);
      } else {
        await createContact(selectedContact);
      }
      toast.success("Contact created successfully");
      getAllContacts();
      setShowModal(false);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-8 h-[calc(100vh-100px)] overflow-y-auto">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Contact</h1>
          <div className="flex gap-2 items-center">
            <Button
              pill
              onClick={() => {
                fileRef.current?.click();
              }}
              color="gray"
            >
              <PiFileXls className="mr-2" /> Import Contacts
            </Button>
            <Button
              gradientDuoTone="purpleToBlue"
              pill
              onClick={() => {
                setShowModal(true);
                setSelectedContact({
                  name: "",
                  is_customer: true,
                  is_supplier: false,
                  is_vendor: false,
                });
              }}
            >
              + Create new Contact
            </Button>
            <div className="cursor-pointer ">
              <LuFilter onClick={() => setDrawerFilter(true)} />
            </div>
          </div>
        </div>
        <Table>
          <Table.Head>
            <Table.HeadCell>Name</Table.HeadCell>
            <Table.HeadCell className="w-48">Email</Table.HeadCell>
            <Table.HeadCell>Phone</Table.HeadCell>
            <Table.HeadCell className="w-64">Address</Table.HeadCell>
            <Table.HeadCell>Product</Table.HeadCell>
            <Table.HeadCell></Table.HeadCell>
          </Table.Head>

          <Table.Body className="divide-y">
            {contacts.length === 0 && (
              <Table.Row>
                <Table.Cell colSpan={5} className="text-center">
                  No contacts found.
                </Table.Cell>
              </Table.Row>
            )}
            {contacts.map((contact) => (
              <Table.Row
                key={contact.id}
                className="bg-white dark:border-gray-700 dark:bg-gray-800"
              >
                <Table.Cell
                  className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold"
                  onClick={() => {}}
                >
                  <div className="flex flex-col">
                    {contact.name}
                    {(contact.tags ?? []).length > 0 && (
                      <div className="flex flex-wrap gap-2">
                        {contact.tags?.map((tag) => (
                          <span
                            className="px-2  text-[8pt] font-semibold text-gray-900 bg-gray-100 rounded dark:bg-gray-700 dark:text-gray-100"
                            key={tag.id}
                            style={{
                              color: getContrastColor(tag.color),
                              backgroundColor: tag.color,
                            }}
                          >
                            {tag.name}
                          </span>
                        ))}
                      </div>
                    )}
                  </div>
                </Table.Cell>
                <Table.Cell>{contact.email}</Table.Cell>
                <Table.Cell>{contact.phone}</Table.Cell>
                <Table.Cell>{contact.address}</Table.Cell>
                <Table.Cell>
                  <div className="flex flex-wrap gap-2">
                    {(contact.products ?? []).map((product) => (
                      <div
                        className="px-2 mb-2 text-[10pt] flex items-center  text-gray-900 bg-gray-100 rounded dark:bg-gray-700 dark:text-gray-100 w-fit gap-1"
                        key={product.id}
                      >
                        <BsTag />
                        <span>{product?.display_name}</span>
                      </div>
                    ))}
                  </div>
                </Table.Cell>
                <Table.Cell>
                  <a
                    href="#"
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                    onClick={() => {
                      setSelectedContact(contact);
                      setShowModal(true);
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
                          `Are you sure you want to delete contact ${contact.name}?`
                        )
                      ) {
                        deleteContact(contact?.id!).then(() => {
                          getAllContacts();
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
        
        <Pagination
          className="mt-4"
          currentPage={page}
          totalPages={pagination?.total_pages ?? 0}
          onPageChange={(val) => {
            setPage(val);
          }}
          showIcons
        />
        <small>Total Record: {pagination?.total ?? 0} items</small>
      </div>
      {selectedContact && (
        <Modal show={showModal} onClose={() => setShowModal(false)}>
          <Modal.Header>
            {selectedContact?.id ? "Edit" : "Create"} Contact
          </Modal.Header>
          <Modal.Body>
            <div className="space-y-4 pb-32">
              <div>
                <Label htmlFor="contactName" value="Name" />
                <TextInput
                  id="contactName"
                  name="name"
                  placeholder="Name"
                  required
                  value={selectedContact?.name ?? ""}
                  onChange={(e) =>
                    setSelectedContact({
                      ...selectedContact,
                      name: e.target.value,
                    })
                  }
                />
              </div>
              <div>
                <Label htmlFor="contactEmail" value="Email" />
                <TextInput
                  id="contactEmail"
                  name="email"
                  type="email"
                  placeholder="Email"
                  value={selectedContact?.email ?? ""}
                  onChange={(e) =>
                    setSelectedContact({
                      ...selectedContact,
                      email: e.target.value,
                    })
                  }
                />
              </div>
              <div>
                <Label htmlFor="contactPhone" value="Phone" />
                <TextInput
                  id="contactPhone"
                  name="phone"
                  type="tel"
                  placeholder="Phone"
                  value={selectedContact?.phone ?? ""}
                  onChange={(e) =>
                    setSelectedContact({
                      ...selectedContact,
                      phone: e.target.value,
                    })
                  }
                />
              </div>
              <div>
                <Label htmlFor="contactAddress" value="Address" />
                <Textarea
                  id="contactAddress"
                  name="address"
                  placeholder="Address"
                  value={selectedContact?.address ?? ""}
                  onChange={(e) =>
                    setSelectedContact({
                      ...selectedContact,
                      address: e.target.value,
                    })
                  }
                />
              </div>
              <div>
                <Label htmlFor="contactPersonPosition" value="Position" />
                <TextInput
                  id="contactPersonPosition"
                  name="contact_person_position"
                  type="text"
                  placeholder="Position"
                  value={selectedContact?.contact_person_position ?? ""}
                  onChange={(e) =>
                    setSelectedContact({
                      ...selectedContact,
                      contact_person_position: e.target.value,
                    })
                  }
                />
              </div>
              <div>
                <Label htmlFor="tag" value="Tag" />
                <CreatableSelect
                  id="tag"
                  name="tag"
                  onCreateOption={(e) => {
                    console.log(e);
                    createTag({
                      name: e,
                      color: randomColor({ luminosity: "dark" }),
                    }).then(() => {
                      getAllTags();
                    });
                  }}
                  isMulti={true}
                  options={tags.map((tag) => ({
                    value: tag.id,
                    label: tag.name,
                    color: tag.color,
                  }))}
                  value={(selectedContact?.tags ?? []).map((tag) => ({
                    value: tag.id,
                    label: tag.name,
                    color: tag.color,
                  }))}
                  onChange={(e) => {
                    setSelectedContact({
                      ...selectedContact,
                      tags: e.map((tag) => ({
                        id: tag.value,
                        name: tag.label,
                        color: tag.color,
                      })),
                    });
                  }}
                  formatOptionLabel={(option) => (
                    <div
                      className="w-fit px-2 py-1 rounded-lg"
                      style={{
                        backgroundColor: option.color,
                        color: getContrastColor(option.color),
                      }}
                    >
                      <span>{option.label}</span>
                    </div>
                  )}
                  formatGroupLabel={(option) => (
                    <div
                      className="w-fit px-2 py-1 rounded-lg"
                      style={{ backgroundColor: "white" }}
                    >
                      <span>{option.label}</span>
                    </div>
                  )}
                />
              </div>
              <div>
                <Label htmlFor="product" value="Products" />
                <Select
                  id="product"
                  name="product"
                  isMulti={true}
                  options={products.map((product) => ({
                    value: product.id,
                    label: product.name,
                  }))}
                  value={(selectedContact?.products ?? []).map((product) => ({
                    value: product.id,
                    label: product.name,
                  }))}
                  onChange={(e) => {
                    setSelectedContact({
                      ...selectedContact!,
                      products: e.map((product) => ({
                        id: product.value!,
                        name: product.label!,
                      })),
                    });
                  }}
                />
              </div>
            </div>
          </Modal.Body>
          <Modal.Footer>
            <div className="flex justify-end w-full">
              <Button onClick={handleCreateContact}>
                {selectedContact?.id ? "Edit" : "Create"} Contact
              </Button>
            </div>
          </Modal.Footer>
        </Modal>
      )}
      <input
        type="file"
        className="hidden"
        onChange={handleFileChange}
        ref={fileRef}
        accept=".xlsx, .xls"
      />
      <Drawer
        open={drawerFilter}
        onClose={function (): void {
          setDrawerFilter(false);
        }}
        position="right"
        style={{ width: "400px" }}
      >
        <Drawer.Header>Filter</Drawer.Header>
        <DrawerItems>
          <div className="mt-8">
            <h1 className="font-semibold text-2xl">Filter</h1>
            <div>
              <Label htmlFor="name" value="Filter" />
              <Select
                isMulti
                value={selectedTags.map((tag) => ({
                  value: tag.id,
                  label: tag.name,
                  color: tag.color,
                }))}
                options={tags.map((tag) => ({
                  value: tag.id,
                  label: tag.name,
                  color: tag.color,
                }))}
                onChange={(e) => {
                  setSelectedTags(
                    e.map((tag) => ({
                      id: tag.value,
                      name: tag.label,
                      color: tag.color,
                    }))
                  );
                }}
                formatOptionLabel={(option) => (
                  <div
                    className="w-fit px-2 py-1 rounded-lg"
                    style={{
                      backgroundColor: option.color,
                      color: getContrastColor(option.color),
                    }}
                  >
                    <span>{option.label}</span>
                  </div>
                )}
              />
            </div>
          </div>
        </DrawerItems>
      </Drawer>
    </AdminLayout>
  );
};
export default ContactPage;
