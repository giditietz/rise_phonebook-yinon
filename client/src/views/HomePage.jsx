import { useState } from "react";

import AddForm from "./AddForm/AddForm";
import ContactList from "./ContactList/ContactList";
import HomePageFooter from "./HomePageFooter/HomePageFooter";
import Modal from "react-modal";

import "./home-page.scss";

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
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  return (
    <>
      <Modal
        isOpen={isAddModalOpen}
        onRequestClose={() => setIsAddModalOpen(false)}
        style={customStyles}
      >
        <AddForm />
      </Modal>
      <div className="home-page-container">
        <ContactList page={page} />
        <HomePageFooter
          page={page}
          setPage={setPage}
          onAddClick={() => setIsAddModalOpen(true)}
        />
      </div>
    </>
  );
};

export default HomePage;
