import matplotlib.pyplot as plt
import pandas as pd

B=[	0.37, 0.42, 0.49, 0.56, 0.65, 0.75, 0.87, 1.00, 1.15, 1.33, 1.54, 1.78,
    2.05, 2.37, 2.74, 3.16, 3.65, 4.22, 4.87, 5.62, 6.49, 7.50, 8.66, 10.00,
    11.55, 13.34, 15.40, 17.78, 20.54, 23.71]

def main(fn):
    df = pd.read_csv(fn, header=None,index_col=0)
    ticks =[
        "..%5.2f" % b if a=="" else ( "%.2f ..      " % a if b=="" else "%.2f .. %.2f" % (a,b) )
        for a, b 
        in zip([""]+B, B+[""]) 
    ]
    plt.style.use("ggplot")
    fig, axes = plt.subplots(nrows=2, ncols=2, figsize=(15, 18))
    for i in range(4):
        ax=axes[i%2, i//2]
        df.iloc[i,:].plot(ax=ax, legend=False, kind="barh", width=0.9)
        ax.set_yticks(list(range(len(ticks))),ticks, fontsize=20, fontname='Menlo')
        ax.set_title(df.index[i], fontsize=40)
        ax.set_xlim(0,3500)
    plt.tight_layout()
    plt.savefig('graph.png')

main("data.csv")