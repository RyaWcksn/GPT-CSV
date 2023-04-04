import os

from langchain.document_loaders import CSVLoader
from langchain.chains import RetrievalQA
from langchain.indexes import VectorstoreIndexCreator
from langchain.llms import OpenAI

os.environ["OPENAI_API_KEY"] = os.environ.get("OPENAI_API_KEY", "")

# Load the document
loader = CSVLoader('./dataCsv/dataset1.csv')

# Create an index using the loaded documents
index_creator = VectorstoreIndexCreator()
docsearch = index_creator.from_loaders([loader])

chain = RetrievalQA.from_chain_type(llm=OpenAI(), chain_type="stuff", retriever=docsearch.vectorstore.as_retriever(), input_key="question")

# Pass a query to the chain
query = "Describe data with Ticket Id of 1206?"
response = chain({"question": query})
print(response['result'])

# Pass a query to the chain
# query = "What is the requested date for data with id 2249?"
# response = chain({"question": query})
# print(response['result'])

