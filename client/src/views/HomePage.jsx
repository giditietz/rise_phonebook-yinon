import { useEffect, useState } from "react";

import ContactForm from "./ContactForm/ContactForm";
import ContactList from "./ContactList/ContactList";
import HomePageFooter from "./HomePageFooter/HomePageFooter";
import Modal from "react-modal";

import "./home-page.scss";
import httpRequest from "../utils/httpRequest/httpRequest";
import SearchRow from "./SearchRow/SearchRow";

const customStyles = {
  content: {
    position: "absolute",
    inset: "14%",
    border: "1px solid rgb(204, 204, 204)",
    background: "rgb(255, 255, 255)",
    overflow: "auto",
    borderRadius: "20px",
    outline: "none",
    padding: "20px",
  },
};

const HomePage = () => {
  const [page, setPage] = useState(1);
  const [isContactFormOpen, setIsContactFormOpen] = useState(false);
  const [contactNum, setContactNum] = useState(0);
  const [lastNameSearchParam, setLastNameSearchParam] = useState("");
  const [firstNameSearchParam, setFirstNameSearchParam] = useState("");
  const [data, setData] = useState([]);
  const [isShow, setIsShow] = useState(false);
  const [contact, setContact] = useState(undefined);
  const [isNew, setIsNew] = useState(false);
  const [isEdit, setIsEdit] = useState(false);

  const resultPerPage = 10;

  const getData = () => {
    httpRequest
      .get(
        `/contacts?page=${
          page - 1
        }&first_name=${firstNameSearchParam}&last_name=${lastNameSearchParam}`
      )
      .then((res) => setData(res.data));
    httpRequest
      .get(`/contacts/contact-num`)
      .then((res) => setContactNum(res.data));
  };

  useEffect(() => {
    getData();
  }, [page, firstNameSearchParam, lastNameSearchParam, isNew]);

  const getNumOfPage = () => {
    return Math.ceil(contactNum / resultPerPage);
  };

  const handleEdit = (item) => {
    setContact(item);
    setIsEdit(true);
    setIsContactFormOpen(true);
  };

  const handleModalClose = () => {
    setIsContactFormOpen(false);
    setIsNew(false);
    setIsShow(false);
    setIsEdit(false);
    setContact({});
  };

  const handleSubmit = () => {
    setIsContactFormOpen(false);
    setIsNew(false);
    getData();
  };

  const handleAddContact = () => {
    setIsContactFormOpen(true);
    setIsNew(true);
  };

  const handleShow = (item) => {
    setIsContactFormOpen(true);
    setIsShow(true);
    setContact(item);
  };

  return (
    <>
      <Modal
        isOpen={isContactFormOpen}
        onRequestClose={() => handleModalClose()}
        style={customStyles}
      >
        <ContactForm
          contact={contact}
          isShow={isShow}
          isNew={isNew}
          isEdit={isEdit}
          handleSubmit={() => handleSubmit()}
        />
      </Modal>
      <div className="home-page-container">
        <SearchRow
          firstName={firstNameSearchParam}
          setFirstName={setFirstNameSearchParam}
          lastName={lastNameSearchParam}
          setLastName={setLastNameSearchParam}
        />
        <ContactList
          onShow={handleShow}
          data={data}
          getData={() => getData()}
          onEdit={handleEdit}
        />
        <HomePageFooter
          numOfPages={getNumOfPage()}
          page={page}
          setPage={setPage}
          onAddClick={() => handleAddContact()}
        />
      </div>
    </>
  );
};

export default HomePage;
