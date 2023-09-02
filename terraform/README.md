# Terraform Wallet Backend

Script Terraform do aplicativo Wallet Backend

## Introdução

Cria ambiente AWS com a estrutura necessária para publicação do aplicativo

### Requisitos

- [Terraform](https://www.terraform.io/downloads.html)
- [aws-cli](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

### Deployment

Prepare o diretório do projeto para o uso do Terraform

```bash
#!/bin/bash

terraform init
```

Quaisquer alterações realizadas no projeto deve ser analisadas antes de sua publicação:

```bash
#!/bin/bash

# compara estrutura local com o publicado
terraform plan
```

Após a validação das modificações, deve-se disparar a sua sincronização com o ambiente de produção

```bash
#!/bin/bash

terraform apply
```

### Rollback

É possível desfazer as alterações aplicadas pelo terraform no provider a partir do seguinte:

```bash
#!/bin/bash

terraform destroy
```

### Contribuições

Contribuições, correções e sugestões de melhoria são muito bem-vindas.