import { api } from '~/common/api';
import { <%= entityUpCase %>Model } from '~/models/<%= entityLowerCase %>.model';
import { TableModel } from '~/models/table.model';

export class <%= entityUpCase %>Service {
  async list(request: Request): Promise<TableModel<<%= entityUpCase %>Model>> {
    const url = new URL(request.url);
    const searchParams = new URLSearchParams(url.search);

    // Padrões podem ser definidos diretamente no URLSearchParams
    const defaultParams = {
      page: '0',
      page_size: '10',
      order_by: '',
      sort_order: '',
      q: '',
    };

    // Atualiza os parâmetros de pesquisa com os valores fornecidos na URL, mantendo os padrões quando não especificados
    Object.entries(defaultParams).forEach(([key, defaultValue]) => {
      const value = searchParams.get(key) || defaultValue;
      if (value) {
        searchParams.set(key, value);
      }
    });

    // Construir a query de forma limpa
    const query = searchParams.toString();

    const response = await api.get(`<%= entityLowerCase %>/?${query}`, {
      request,
    });

    return new TableModel<<%= entityUpCase %>Model>(<%= entityUpCase %>Model, response);
  }

  async createOrEdit({ body }: { body: any }, request: Request): Promise<void> {
    const payload = new <%= entityUpCase %>Model(body);

    if (payload.id && payload.id > 0) {
      await api.put(`<%= entityLowerCase %>/${payload.id}`, {body: payload.toJson(), request});
    } else {
      await api.post('<%= entityLowerCase %>', {body: payload.toJson(), request});
    }
  }

  async update<%= entityUpCase %>(body : any, request: Request): Promise<void> {

    const payload = new <%= entityUpCase %>Model(body);
    await api.put(`<%= entityLowerCase %>/${body.id}`, {body: payload, request});
  }
}
