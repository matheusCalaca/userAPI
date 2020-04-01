import React from 'react';
import { Form, Scope } from '@unform/core';
import Input from '../../components/form/Input';





export default function CadastroPessoa() {
    function handleSubmit(data) {
        console.log(data);
        // { email: 'test@example.com', password: '123456' }
    }
    return (
        <>
            <Form onSubmit={handleSubmit}>
                <Input name="nome" type="text" />
                <Input name="sebrenome" type="text" />
                <Input name="CPF" type="text" />
                <Input name="RG" type="text" />
                <Input name="dataNascimento" type="date" />
                <Input name="email" type="text" />
                <Scope path="telefone[0]">
                    <Input name="ddd" type="number"/>
                    <Input name="numero" type="text" />
                    <Input name="tipo" type="text" />
                </Scope>
                <button type="submit">Cadastrar</button>
            </Form>
        </>
    );
}