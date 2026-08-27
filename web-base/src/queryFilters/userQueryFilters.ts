import type { QueryParameter } from "@/api";

export class UserQueryFilters {
    static WithUsernames(usernames: string[]): QueryParameter {
        return {
            Name: "withusernames",
            Value: usernames
        };
    }

    static WithIds(ids: string[]): QueryParameter {
        return {
            Name: "withids",
            Value: ids
        };
    }

    static IsActive(isActive: boolean): QueryParameter {
        return {
            Name: "isactive",
            Value: [isActive.toString()]
        };
    }

    static Empty(): QueryParameter {
        return {
            Name: "empty",
            Value: []
        };
    }
}