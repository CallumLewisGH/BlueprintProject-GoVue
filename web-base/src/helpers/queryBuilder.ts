import type { QueryParameter } from "@/api";

export class QueryBuilder {
    private parameters: QueryParameter[] = [];

    addParameter(parameter: QueryParameter): QueryBuilder {
        this.parameters.push(parameter);
        return this;
    }

    build(): QueryParameter[] {
        return this.parameters;
    }

}

