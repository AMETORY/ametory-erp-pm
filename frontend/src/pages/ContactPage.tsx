import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { PaginationResponse } from "../objects/pagination";
import {
  createContact,
  deleteContact,
  getContacts,
  updateContact,
} from "../services/api/contactApi";
import { ContactModel } from "../models/contact";
import { LoadingContext } from "../contexts/LoadingContext";
import { getPagination } from "../utils/helper";
import toast from "react-hot-toast";
import {
  Button,
  Label,
  Modal,
  Pagination,
  Table,
  Textarea,
  TextInput,
} from "flowbite-react";

interface ContactPageProps {}

const ContactPage: FC<ContactPageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [mounted, setMounted] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [contacts, setContacts] = useState<ContactModel[]>([]);
  const [showModal, setShowModal] = useState(false);
  const [selectedContact, setSelectedContact] = useState<ContactModel>();

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllContacts();
    }
  }, [mounted]);

  const getAllContacts = async () => {
    try {
      setLoading(true);
      let resp: any = await getContacts({ page, size, search });
      setContacts(resp.data.items);
      setPagination(getPagination(resp.data));
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
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Contact</h1>
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
        </div>
        <Table>
          <Table.Head>
            <Table.HeadCell>Name</Table.HeadCell>
            <Table.HeadCell>Email</Table.HeadCell>
            <Table.HeadCell>Phone</Table.HeadCell>
            <Table.HeadCell>Address</Table.HeadCell>
            <Table.HeadCell>Position</Table.HeadCell>
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
                  {contact.name}
                </Table.Cell>
                <Table.Cell>{contact.email}</Table.Cell>
                <Table.Cell>{contact.phone}</Table.Cell>
                <Table.Cell>{contact.address}</Table.Cell>
                <Table.Cell>{contact.contact_person_position}</Table.Cell>
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
      </div>
      {selectedContact && (
        <Modal show={showModal} onClose={() => setShowModal(false)}>
          <Modal.Header>{selectedContact?.id ? "Edit" : "Create"} Contact</Modal.Header>
          <Modal.Body>
            <div className="space-y-4">
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
    </AdminLayout>
  );
};
export default ContactPage;
