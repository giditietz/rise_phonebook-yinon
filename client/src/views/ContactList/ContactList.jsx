import { useState, useEffect } from "react";

import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";

import httpRequest from "../../utils/httpRequest/httpRequest";

import "./contact-list.scss";

import text from "../../utils/language/text.json";

const ContactList = () => {
  const [data, setData] = useState([]);
  const [page, setPage] = useState(0);

  useEffect(() => {
    httpRequest.get(`/contacts?page=${page}`).then((res) => setData(res.data));
  }, [page]);

  const getMaxContactItem = (item) => {
    return item?.phone?.length < item?.address?.length
      ? item?.address?.length
      : item?.phone?.length;
  };

  const getColSpan = {
    "First name": 1,
    "Last name": 1,
    "Phone number": 2,
    Address: 5,
  };
  const getRowSpan = {
    "First name": 2,
    "Last name": 2,
    "Phone number": 1,
    Address: 1,
  };

  const mergeSubTitle = [...text.phoneSubTitle, ...text.addressSubTitle];
  return (
    <div className="table-container">
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              {text.gridTitle.map((item) => {
                return (
                  <TableCell
                    align="center"
                    sx={{ border: "1px solid #c4c4c4" }}
                    key={item}
                    colSpan={getColSpan[item]}
                    rowSpan={getRowSpan[item]}
                  >
                    {item}
                  </TableCell>
                );
              })}
            </TableRow>
            <TableRow>
              {mergeSubTitle.map((item, index) => {
                return (
                  <TableCell
                    align="center"
                    sx={{ border: "1px solid #c4c4c4" }}
                    key={`${item}${index}`}
                  >
                    {item}
                  </TableCell>
                );
              })}
            </TableRow>
          </TableHead>
          <TableBody>
            {data.map((item) => (
              <TableRow
                key={item.contactID}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
              >
                <TableCell
                  component="th"
                  scope="row"
                  rowSpan={getMaxContactItem(item)}
                >
                  {item.firstName}
                </TableCell>
                <TableCell align="right" rowSpan={getMaxContactItem(item)}>{item.lastName}</TableCell>
                {item?.phone?.length
                  ? item.phone.map((phone) => {
                      <>
                        <TableCell align="right">{phone.description}</TableCell>
                        <TableCell align="right">
                          {phone.phone_number}
                        </TableCell>
                      </>;
                    })
                  : <><TableCell align="right"></TableCell>
                  <TableCell align="right"></TableCell>
                  </>}
                {/* {item.address ? (
                <TableCell align="right">{item.address[0]}</TableCell>
              ) : null} */}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
};

export default ContactList;
