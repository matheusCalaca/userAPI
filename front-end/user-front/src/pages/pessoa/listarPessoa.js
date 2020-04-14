import React, { Component } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Paper, TableRow, TablePagination, TableHead, Table, TableBody, TableCell, TableContainer } from '@material-ui/core';
import Axios from 'axios';


const columns = [
    { id: 'nome', label: 'Nome', minWidth: 50 },
    { id: 'CPF', label: 'CPF', minWidth: 16 },
    {
        id: 'email',
        label: 'E-mail',
        minWidth: 100,
        format: (value) => value.toLocaleString(),
    },
];

function createData(nome, CPF, email) {
    return { nome, CPF, email };
}

const useStyles = makeStyles({
    root: {
        width: '100%',
    },
    container: {
        maxHeight: 440,
    },
});

class ListaPessoa extends Component {

    handleChangePage = (event, newPage) => {
        this.state.page = newPage;
    };

    handleChangeRowsPerPage = (event) => {
        this.state.rowsPerPage += event.target.value;
        this.state.page = 0;
    };
    constructor(props) {
        super(props);

        this.state = {
            classes: useStyles,
            page: 0,
            rows: [],
            setPage: 0,
            rowsPerPage: 10,
            setRowsPerPage: 10
        };
    }

    componentDidMount() {
        this.returnListPessoas();
    }
    //   returnListPessoas retorna as pessoa do servidor
    returnListPessoas = async () => Axios.get(`http://localhost:8000/api/users`).then(
        resp => {
            this.state.rows = []
            var pesssoas = JSON.parse(JSON.stringify(resp.data));
            for (let i = 0; i < pesssoas.length; i++) {
                const pessoa = pesssoas[i];
                this.state.rows.push(createData(pessoa['nome'], pessoa['CPF'], pessoa['email']));
            }
            this.setState({ test: this.state.rows })
        }
    ).catch(
        err => {
            console.log("ERRO - " + JSON.stringify(err));
            console.log(err);
        }
    );

    render() {

        return (
            <Paper className={this.state.classes.root}>
                <TableContainer className={this.state.classes.container}>
                    <Table stickyHeader aria-label="sticky table">
                        <TableHead>
                            <TableRow>
                                {columns.map((column) => (
                                    <TableCell
                                        key={column.id}
                                        align={column.align}
                                        style={{ minWidth: column.minWidth }}
                                    >
                                        {column.label}
                                    </TableCell>
                                ))}
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.state.rows.slice(this.state.page * this.state.rowsPerPage, this.state.page * this.state.rowsPerPage + this.state.rowsPerPage).map((row) => {
                                return (
                                    <TableRow hover role="checkbox" tabIndex={-1} key={row.CPF}>
                                        {columns.map((column) => {
                                            const value = row[column.id];
                                            return (
                                                <TableCell key={column.id} align={column.align}>
                                                    {column.format && typeof value === 'number' ? column.format(value) : value}
                                                </TableCell>
                                            );
                                        })}
                                    </TableRow>
                                );
                            })}
                        </TableBody>
                    </Table>
                </TableContainer>
                <TablePagination
                    rowsPerPageOptions={[10, 25, 100]}
                    component="div"
                    count={this.state.rows.length}
                    rowsPerPage={this.state.rowsPerPage}
                    page={this.state.page}
                    onChangePage={this.handleChangePage}
                    onChangeRowsPerPage={this.handleChangeRowsPerPage}
                />
            </Paper>
        );
    }
}
export default ListaPessoa; 