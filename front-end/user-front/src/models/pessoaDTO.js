import { TelefoneDTO } from "./telefoneDTO";
import { EnderecoDTO } from "./enderecoDTO";

export class  PessoaDto   {
     CPF;
     nome;
     sobrenome;
     RG;
     dataNascimento = new Date();
     email;
     telefones = [new TelefoneDTO()];
     enderecos = [new EnderecoDTO()];


}