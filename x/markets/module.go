// Markets module is used to store Base-Quote pairs with specific settings.
// Order can't be posted without a corresponding market object.
package markets

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/dfinance/dnode/x/markets/client"
	"github.com/dfinance/dnode/x/markets/client/rest"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic app module basics object.
type AppModuleBasic struct{}

var _ module.AppModuleBasic = AppModuleBasic{}

// Name gets module name.
func (AppModuleBasic) Name() string {
	return ModuleName
}

// RegisterCodec registers module codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// DefaultGenesis gets default module genesis state.
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

// ValidateGenesis validates module genesis state.
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	state := GenesisState{}
	ModuleCdc.MustUnmarshalJSON(bz, &state)

	return state.Validate()
}

// RegisterRESTRoutes registers module REST routes.
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

// GetTxCmd returns module root tx command.
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return client.GetTxCmd(cdc)
}

// GetQueryCmd returns module root query command.
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return client.GetQueryCmd(cdc)
}

// AppModule is a app module type.
type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

// NewAppModule creates new AppModule object.
func NewAppModule(keeper Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	}
}

// Name gets module name.
func (app AppModule) Name() string {
	return ModuleName
}

// RegisterInvariants registers module invariants.
func (app AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns module messages route.
func (app AppModule) Route() string {
	return ModuleName
}

// NewHandler returns module messages handler.
func (app AppModule) NewHandler() sdk.Handler {
	return NewHandler(app.keeper)
}

// QuerierRoute returns module querier route.
func (app AppModule) QuerierRoute() string {
	return ModuleName
}

// NewQuerierHandler creates module querier.
func (app AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(app.keeper)
}

// InitGenesis inits module-genesis state.
func (app AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	app.keeper.InitGenesis(ctx, data)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports module genesis state.
func (app AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return app.keeper.ExportGenesis(ctx)
}

// BeginBlock performs module actions at a block start.
func (app AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock performs module actions at a block end.
// It returns no validator updates.
func (app AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
