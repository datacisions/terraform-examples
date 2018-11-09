## Basics

- EC2 Instance
- "testpair"
- not secure, allows SSH access from everywhere currently!

### Loading data into Neptune

Load data not into the cluster endpoint, but into the Writer:

```
curl -X POST -d "{\"gremlin\":\"g.addV('supplier').property('name','supplier1').property('date','10.02.2011').addV('consumer').property('name','consumer1').property('date','05.08.2000')\"}" http://{your-neptune-endpoint}:8182/gremlin

curl -X POST -d "{\"gremlin\":\"g.V().has('supplier','name','supplier1').as('a').V().has('consumer','name','consumer1').addE('deliversTo').to('a')\"}"  http://{your-neptune-endpoint}:8182/gremlin
```

Now query (the cluster) to get some data.

```
curl -X POST --data-binary 'query=select ?s ?p ?o where {?s ?p ?o} limit 10'\
http://{your-neptune-endpoint}:8182/sparql

curl -X POST -d '{"gremlin":"g.V().limit(1)"}' http://{your-neptune-endpoint}:8182/gremlin

>>
{"requestId":"22b33390-77fc-be22-b673-02c9171a857f","status":{"message":"","code":200,"attributes":{"@type":"g:Map","@value":[]}},"result":{"data":{"@type":"g:List","@value":[{"@type":"g:Vertex","@value":{"id":"ceb3338f-7175-077d-ecec-68aaafcc83b7","label":"supplier","properties":{"date":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-1293060726},"value":"10.02.2011","label":"date"}}],"name":[{"@type":"g:VertexProperty","@value":{"id":{"@type":"g:Int32","@value":-2023617795},"value":"supplier1","label":"name"}}]}}}]},"meta":{"@type":"g:Map","@value":[]}}}
```

## More Resources

https://github.com/aws-samples/amazon-neptune-samples/tree/master/gremlin/collaborative-filtering
