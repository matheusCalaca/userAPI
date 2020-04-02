import React, { Component } from 'react';
import { Form } from '@unform/web';
import { Scope } from '@unform/core'
import Input from '../../components/form/Input';
import PhoneForm from '../../components/form/phone-form'




class CadastroPessoa extends Component {
     handleSubmit = (data) => {
        console.log(JSON.stringify(data));
        // { email: 'test@example.com', password: '123456' }
    }
    state = {
        phoneContador: 1,
    }

    handlePlusTelefone = () => {
        this.setState({
            phoneContador: this.state.phoneContador + 1
        })
    }

render(){
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
            <PhoneForm teste={this.state.phoneContador}></PhoneForm>
            <button onClick={this.handlePlusTelefone}>adicionar</button><br /><br />

            <Scope path="enderecos[0]">
                <span>CEP: </span> <br />
                <Input name="CEP" type="number" /> <br />
                <span>logradouro: </span> <br />
                <Input name="logradouro" type="text" /> <br />
                <span>bairro: </span> <br />
                <Input name="bairro" type="text" /> <br />
                <span>cidade: </span> <br />
                <Input name="cidade" type="text" /> <br />
                <span>UF: </span> <br />
                <Input name="UF" type="text" /> <br />
                <span>complemento: </span> <br />
                <Input name="complemento" type="text" /> <br />
                <span>numero: </span> <br />
                <Input name="numero" type="text" /> <br />
                <span>tipo: </span> <br />
                <Input name="tipo" type="text" /> <br />
            </Scope>
            <button type="submit">Cadastrar</button>
        </Form>
    );
        }
}

export default CadastroPessoa;