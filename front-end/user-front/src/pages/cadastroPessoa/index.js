import React, { Component } from 'react';
import { Form } from '@unform/web';
import Input from '../../components/form/Input';
import PhoneForm from '../../components/form/phone-form'
import EnderecoForm from '../../components/form/endereco-form';




class CadastroPessoa extends Component {
    handleSubmit = (data) => {
        console.log(JSON.stringify(data));
    }

    render() {
        return (
            <Form onSubmit={this.handleSubmit}>
                <span>nome: </span> <br />
                <Input name="nome" type="text" /> <br />
                <span>sebrenome: </span> <br />
                <Input name="sobrenome" type="text" /> <br />
                <span>CPF: </span> <br />
                <Input name="CPF" type="text" /> <br />
                <span>RG: </span> <br />
                <Input name="RG" type="text" /> <br />
                <span>dataNascimento: </span> <br />
                <Input name="dataNascimento" type="date" /> <br />
                <span>email: </span> <br />
                <Input name="email" type="email" /> <br />
                <PhoneForm ></PhoneForm>
                <EnderecoForm></EnderecoForm>

                <button type="submit">Cadastrar</button>
            </Form>
        );
    }
}

export default CadastroPessoa;