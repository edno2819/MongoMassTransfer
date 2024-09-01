### Pacotes:
    Para criar um pacote vc deve criar uma pasta com o mesmo nome do pacote
    Os arquivos dentro desta pasta deve ter: package nome_da_pasta
    A importação deve ter o nome_do_projeto + path até o pacote ex: src/utils
    Apenas o que começar com a letra maiúscula será EXPORTADO

### todo:
    - Pra que server "call" : Usado para chamar funções fora do contexto. ex: debug

// Lembrar de usar o debug com f5


- interface{} é para definir o tipo genérico de struc


- Referências para criar um leitor
https://kgrz.io/reading-files-in-go-an-overview.html

## Quantity of a chunk insert on Mongodb
A quantidade ideal de documentos a serem inseridos em uma única operação insert_many no MongoDB pode variar dependendo de vários fatores, como a capacidade do hardware, o tamanho dos documentos, a configuração do cluster MongoDB, e a latência de rede.

No entanto, com base em recomendações gerais e práticas comuns:

Batch Size: Recomenda-se que cada lote (batch) de inserções via insert_many seja de aproximadamente 1.000 a 5.000 documentos. Esse tamanho geralmente proporciona um bom equilíbrio entre a carga de trabalho no servidor e a eficiência da operação. Lotes menores podem causar overhead adicional, enquanto lotes muito grandes podem exigir mais memória e processamento, impactando negativamente a performance.

Tamanho Total dos Dados: O MongoDB impõe um limite de 16 MB por lote de inserção. Portanto, mesmo que você consiga inserir 5.000 documentos em um único lote, deve garantir que o tamanho total dos documentos não ultrapasse 16 MB. Se o lote exceder esse limite, o MongoDB retornará um erro.

Performance Tuning: Ajustar o tamanho do lote para otimizar a performance depende do seu ambiente específico. Testes A/B e monitoramento de desempenho são cruciais para encontrar o tamanho ideal de lote para seu cenário.

Portanto, um bom ponto de partida seria configurar seus lotes entre 1.000 a 5.000 documentos, assegurando que o tamanho total não ultrapasse 16 MB. A partir daí, você pode experimentar e ajustar conforme necessário para otimizar ainda mais o desempenho no seu ambiente específico.