/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';

export default function useTfchain4Tfchain4() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.Tfchain4Tfchain4.query.queryParams().then( res => res.data );
    }, options);
  }
  
  return {QueryParams,
  }
}
