import pandas as pd
from sklearn.tree import DecisionTreeRegressor

melbourne_file_path = '/Users/antonkapralov/dev/data/melb_data.csv'

# read the data and store data in DataFrame titled melbourne_data
melbourne_data = pd.read_csv(melbourne_file_path)
# print a summary of the data in Melbourne data
print(melbourne_data[['Price']].head())

# print(melbourne_data.columns)

X = melbourne_data[['Rooms', 'Bathroom', 'Landsize', 'Lattitude', 'Longtitude']]
y = melbourne_data.Price

# Define model. Specify a number for random_state to ensure same results each run
melbourne_model = DecisionTreeRegressor(max_leaf_nodes=500, random_state=1).fit(X, y)

print(X.head())

x = pd.DataFrame({
    'Rooms': [3],
    'Bathroom': [2],
    'Landsize': [202],
    'Lattitude': [-37.7996],
    'Longtitude': [144.9984]
})
print(melbourne_model.predict(x))
